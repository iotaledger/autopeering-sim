package simulation

import (
	"time"

	"github.com/iotaledger/hive.go/autopeering/peer"
	"github.com/iotaledger/hive.go/autopeering/peer/service"
	"github.com/iotaledger/hive.go/autopeering/salt"
	"github.com/iotaledger/hive.go/autopeering/selection"
	"github.com/iotaledger/hive.go/autopeering/server"
	"github.com/iotaledger/hive.go/autopeering/transport"
	"github.com/iotaledger/hive.go/database/mapdb"
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
	services := service.New()
	services.Update(service.PeeringKey, trans.LocalAddr().Network(), trans.LocalAddr().String())
	db, _ := peer.NewDB(mapdb.NewMapDB())

	local, _ := peer.NewLocal(services, db)

	s, _ := salt.NewSalt(saltLifetime)
	local.SetPrivateSalt(s)
	s, _ = salt.NewSalt(saltLifetime)
	local.SetPublicSalt(s)

	prot := selection.New(local, discover, selection.Logger(log), selection.DropOnUpdate(dropOnUpdate))
	srv := server.Serve(local, network.GetTransport(name), log, prot)

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
