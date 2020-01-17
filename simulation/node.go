package simulation

import (
	"time"

	"github.com/iotaledger/goshimmer/packages/autopeering/peer"
	"github.com/iotaledger/goshimmer/packages/autopeering/salt"
	"github.com/iotaledger/goshimmer/packages/autopeering/selection"
	"github.com/iotaledger/goshimmer/packages/autopeering/server"
	"github.com/iotaledger/goshimmer/packages/autopeering/transport"
	"github.com/iotaledger/hive.go/logger"
)

type Node struct {
	local *peer.Local
	prot  *selection.Protocol

	Start func()
	Stop  func()
}

func NewNode(name string, saltLifetime time.Duration, network *transport.ChanNetwork, dropOnUpdate bool, discover selection.DiscoverProtocol) Node {
	log := logger.NewLogger(name)

	network.AddTransport(name)
	trans := network.GetTransport(name)
	db := peer.NewMemoryDB(log)
	local, _ := peer.NewLocal(trans.LocalAddr().Network(), trans.LocalAddr().String(), db)

	s, _ := salt.NewSalt(saltLifetime)
	local.SetPrivateSalt(s)
	s, _ = salt.NewSalt(saltLifetime)
	local.SetPublicSalt(s)

	cfg := selection.Config{Log: log, DropOnUpdate: dropOnUpdate}
	prot := selection.New(local, discover, cfg)
	srv := server.Listen(local, network.GetTransport(name), log, prot)

	return Node{
		local: local,
		prot:  prot,
		Start: func() {
			prot.Start(srv)
		},
		Stop: func() {
			prot.Close()
			srv.Close()
			trans.Close()
			db.Close()
		},
	}
}

func (n Node) ID() peer.ID {
	return n.local.ID()
}

func (n Node) Peer() *peer.Peer {
	return &n.local.Peer
}

func (n Node) GetNeighbors() []*peer.Peer {
	return n.prot.GetNeighbors()
}

func (n Node) GetOutgoingNeighbors() []*peer.Peer {
	return n.prot.GetOutgoingNeighbors()
}
