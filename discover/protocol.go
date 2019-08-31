package discover

import (
	"bytes"
	"time"

	"github.com/wollac/autopeering/peer"
	peerpb "github.com/wollac/autopeering/peer/proto"
	pb "github.com/wollac/autopeering/proto"
)

// ------ message senders ------

// ping sends a ping to the specified peer and blocks until a reply is received
// or the packe timeout.
func (s *Server) ping(to *peer.Peer) error {
	return <-s.sendPing(to.ID(), to.Address())
}

// sendPing sends a ping to the specified address and expects a matching reply.
// This method is non-blocking, but it returns a channel that can be used to query potential errors.
func (s *Server) sendPing(toID peer.ID, toAddr string) <-chan error {
	ping := newPing(s.LocalAddr(), toAddr)
	pkt := s.encode(ping)
	// compute the message hash
	hash := packetHash(pkt.GetData())

	// Add a matcher for the reply to the pending reply queue. Pongs are matched,
	// if they come from the specified peer and reference the ping we're about to send.
	errc := s.expectReply(toAddr, toID, pb.MPong, func(m pb.Message) (bool, bool) {
		matched := bytes.Equal(m.(*pb.Pong).GetPingHash(), hash)
		return matched, matched
	})

	// send the ping
	s.write(toAddr, ping.Name(), pkt)
	return errc
}

// requestPeers request known peers from the given target. This method blocks
// until a response is received and the provided peers are returned.
func (s *Server) requestPeers(to *peer.Peer) ([]*peer.Peer, error) {
	toID := to.ID()
	toAddr := to.Address()
	s.ensureVerified(toID, toAddr)

	// create the request package
	req := newPeersRequest()
	pkt := s.encode(req)
	// compute the message hash
	hash := packetHash(pkt.GetData())
	peers := make([]*peer.Peer, 0, maxPeersInResponse)

	errc := s.expectReply(toAddr, to.ID(), pb.MPeersResponse, func(m pb.Message) (bool, bool) {
		res := m.(*pb.PeersResponse)
		if !bytes.Equal(res.GetReqHash(), hash) {
			return false, false
		}

		for _, rp := range res.GetPeers() {
			p, err := peer.FromProto(rp)
			if err != nil {
				s.log.Warnw("invalid peer received", "err", err)
				continue
			}
			peers = append(peers, p)
		}

		return true, true
	})

	// send the request and wait for the response
	s.write(toAddr, req.Name(), pkt)
	return peers, <-errc
}

// ------ helper functions ------

// isVerified checks whether the given peer has recently been verified.a recent enough endpoint proof.
func (s *Server) isVerified(id peer.ID, address string) bool {
	return time.Since(s.mgr.db.LastPong(id, address)) < pongExpiration
}

// ensureVerified checks if the given peer has recently sent a ping;
// if not, we send a ping to trigger a verification.
func (s *Server) ensureVerified(id peer.ID, address string) {
	if time.Since(s.mgr.db.LastPing(id, address)) >= pongExpiration {
		<-s.sendPing(id, address)
		// Wait for them to ping back and process our pong
		time.Sleep(responseTimeout)
	}
}

// expired checks whether the given UNIX time stamp is too far in the past.
func expired(ts int64) bool {
	return time.Since(time.Unix(ts, 0)) >= packetExpiration
}

// ------ Packet Constructors ------

func newPing(fromAddr string, toAddr string) *pb.Ping {
	return &pb.Ping{
		Version:   VersionNum,
		From:      fromAddr,
		To:        toAddr,
		Timestamp: time.Now().Unix(),
	}
}

func newPong(toAddr string, reqData []byte) *pb.Pong {
	return &pb.Pong{
		PingHash: packetHash(reqData),
		To:       toAddr,
	}
}

func newPeersRequest() *pb.PeersRequest {
	return &pb.PeersRequest{
		Timestamp: time.Now().Unix(),
	}
}

func newPeersResponse(reqData []byte, list []*peer.Peer) *pb.PeersResponse {
	peers := make([]*peerpb.Peer, 0, len(list))
	for _, p := range list {
		peers = append(peers, p.ToProto())
	}
	return &pb.PeersResponse{
		ReqHash: packetHash(reqData),
		Peers:   peers,
	}
}

// ------ Packet Handlers ------

func (s *Server) validatePing(m *pb.Ping, fromAddr string) bool {
	// check version number
	if m.GetVersion() != VersionNum {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"version", m.GetVersion(),
		)
		return false
	}
	// check that From matches the package sender address
	if m.GetFrom() != fromAddr {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"from", m.GetFrom(),
		)
		return false
	}
	// check that To matches the local address
	if m.GetTo() != s.LocalAddr() {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"to", m.GetTo(),
		)
		return false
	}
	// check Timestamp
	if expired(m.GetTimestamp()) {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"ts", time.Unix(m.GetTimestamp(), 0),
		)
		return false
	}
	return true
}

func (s *Server) handlePing(m *pb.Ping, fromID peer.ID, fromAddr string, rawData []byte) {
	s.mgr.db.UpdateLastPing(fromID, fromAddr, time.Now())

	// create and send the pong response
	pong := newPong(fromAddr, rawData)
	s.send(fromAddr, pong)

	// if the peer is new or expired, send a ping to verify
	if !s.isVerified(fromID, fromAddr) {
		s.sendPing(fromID, fromAddr)
	}
}

func (s *Server) validatePong(m *pb.Pong, fromID peer.ID, fromAddr string) bool {
	// check that To matches the local address
	if m.GetTo() != s.LocalAddr() {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"to", m.GetTo(),
		)
		return false
	}
	// there must be a ping waiting for this pong as a reply
	if !s.handleReply(fromAddr, fromID, m) {
		s.log.Debugw("no matching request",
			"type", m.Name(),
			"from", fromAddr,
		)
		return false
	}
	return true
}

func (s *Server) handlePong(m *pb.Pong, fromID peer.ID, fromAddr string, fromKey peer.PublicKey) {
	s.mgr.db.UpdateLastPong(fromID, fromAddr, time.Now())
	// a valid pong verifies the peer
	s.mgr.addVerifiedPeer(peer.NewPeer(fromKey, fromAddr))
}

func (s *Server) validatePeersRequest(m *pb.PeersRequest, fromID peer.ID, fromAddr string) bool {
	// check Timestamp
	if expired(m.GetTimestamp()) {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"ts", time.Unix(m.GetTimestamp(), 0),
		)
		return false
	}
	if !s.isVerified(fromID, fromAddr) {
		s.log.Debugw("failed to validate",
			"type", m.Name(),
			"id", fromID,
			"addr", fromAddr,
		)
		return false
	}
	return true
}

func (s *Server) handlePeersRequest(m *pb.PeersRequest, fromID peer.ID, fromAddr string, rawData []byte) {
	// get a random list of verified peers
	peers := s.mgr.getRandomPeers(maxPeersInResponse, 1)
	s.send(fromAddr, newPeersResponse(rawData, peers))
}

func (s *Server) validatePeersResponse(m *pb.PeersResponse, fromID peer.ID, fromAddr string) bool {
	// there must not be too many peers
	if len(m.GetPeers()) > maxPeersInResponse {
		return false
	}
	// there must be a request waiting for this response
	if !s.handleReply(fromAddr, fromID, m) {
		s.log.Debugw("no matching request",
			"type", m.Name(),
			"from", fromID,
		)
		return false
	}
	return true
}
