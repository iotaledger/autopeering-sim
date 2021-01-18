package simulation

import (
	"fmt"
	"net"
	"time"

	"github.com/iotaledger/autopeering-sim/simulation/transport"
	"github.com/iotaledger/hive.go/autopeering/peer"
	"github.com/iotaledger/hive.go/autopeering/peer/service"
	"github.com/iotaledger/hive.go/autopeering/selection"
	"github.com/iotaledger/hive.go/autopeering/server"
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/kvstore/mapdb"
	"github.com/iotaledger/hive.go/logger"
)

type Node struct {
	local *peer.Local
	Prot  *selection.Protocol

	Start func()
	Stop  func()
}

func NewNode(id transport.PeerID, saltLifetime time.Duration, network *transport.Network, dropOnUpdate bool, discover selection.DiscoverProtocol) Node {
	log := logger.NewLogger(fmt.Sprintf("peer%d", id))

	conn, _ := network.Listen(id, 0)

	services := service.New()
	services.Update(service.PeeringKey, conn.LocalAddr().String(), 0)
	db, _ := peer.NewDB(mapdb.NewMapDB())

	local, _ := peer.NewLocal(conn.LocalAddr().(*net.UDPAddr).IP, services, db)

	prot := selection.New(local, discover, selection.Logger(log), selection.DropOnUpdate(dropOnUpdate))
	srv := server.Serve(local, conn, log, prot)

	return Node{
		local: local,
		Prot:  prot,
		Start: func() {
			prot.Start(srv)
		},
		Stop: func() {
			prot.Close()
			srv.Close()
			conn.Close()
		},
	}
}

func (n Node) ID() identity.ID {
	return n.local.ID()
}

func (n Node) Peer() *peer.Peer {
	return n.local.Peer
}

func (n Node) GetNeighbors() []*peer.Peer {
	return n.Prot.GetNeighbors()
}

func (n Node) GetOutgoingNeighbors() []*peer.Peer {
	return n.Prot.GetOutgoingNeighbors()
}
