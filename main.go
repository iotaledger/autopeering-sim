package main

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/iotaledger/autopeering-sim/simulation"
	"github.com/iotaledger/autopeering-sim/simulation/config"
	"github.com/iotaledger/autopeering-sim/simulation/transport"
	"github.com/iotaledger/autopeering-sim/simulation/visualizer"
	"github.com/iotaledger/hive.go/autopeering/peer"
	"github.com/iotaledger/hive.go/autopeering/selection"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/logger"
	"github.com/spf13/viper"
)

var (
	nodeMap map[identity.ID]simulation.Node

	srv     *visualizer.Server
	closing chan struct{}
	wg      sync.WaitGroup
)

// dummyDiscovery is a dummy implementation of DiscoveryProtocol never returning any verified peers.
type dummyDiscovery struct{}

func (d dummyDiscovery) IsVerified(identity.ID, net.IP) bool       { return true }
func (d dummyDiscovery) EnsureVerified(*peer.Peer) error           { return nil }
func (d dummyDiscovery) SendPing(string, identity.ID) <-chan error { panic("not implemented") }
func (d dummyDiscovery) GetVerifiedPeer(id identity.ID) *peer.Peer { return nodeMap[id].Peer() }
func (d dummyDiscovery) GetVerifiedPeers() []*peer.Peer            { return getAllPeers() }

var discover = &dummyDiscovery{}

func getAllPeers() []*peer.Peer {
	result := make([]*peer.Peer, 0, len(nodeMap))
	for _, node := range nodeMap {
		result = append(result, node.Peer())
	}
	return result
}

func runSim() {
	log.Println("Run simulation with the following parameters:")
	config.PrintConfig()

	selection.SetParameters(selection.Parameters{
		SaltLifetime:           config.SaltLifetime(),
		OutboundUpdateInterval: 200 * time.Millisecond, // use exactly the same update time as previously
	})

	closure := events.NewClosure(func(ev *selection.PeeringEvent) {
		if ev.Status {
			log.Printf("Peering: %s<->%s\n", ev.Self.String(), ev.Peer.ID())
		}
	})
	selection.Events.OutgoingPeering.Attach(closure)
	defer selection.Events.OutgoingPeering.Detach(closure)

	//lambda := (float64(N) / SaltLifetime.Seconds()) * 10
	initialSalt := 0.

	log.Println("Creating peers ...")
	netw := transport.NewNetwork()
	nodeMap = make(map[identity.ID]simulation.Node, config.NumberNodes())

	for i := 0; i < config.NumberNodes(); i++ {
		node := simulation.NewNode(
			transport.PeerID(i),
			time.Duration(initialSalt)*time.Second,
			netw,
			config.DropOnUpdate(),
			discover,
			config.R(),
			config.Ro(),
		)
		nodeMap[node.ID()] = node

		if config.VisEnabled() {
			visualizer.AddNode(node.ID().String())
		}

		// initialSalt = initialSalt + (1 / lambda)				 // constant rate
		// initialSalt = initialSalt + rand.ExpFloat64()/lambda  // poisson process
		initialSalt = rand.Float64() * config.SaltLifetime().Seconds() // random
	}
	log.Println("Creating peers ... done")

	log.Println("Populating mana ...")

	identities := []*identity.Identity{}
	for _, node := range nodeMap {
		identities = append(identities, node.Peer().Identity)
	}
	simulation.IdentityMana = simulation.NewZipfMana(identities, 0.82)

	log.Println("Populating mana ... done")

	analysis := simulation.NewLinkAnalysis(nodeMap)

	if config.VisEnabled() {
		statVisualizer()
	}

	log.Println("Starting peering ...")
	analysis.Start()
	for _, node := range nodeMap {
		node.Start()
	}
	log.Println("Starting peering ... done")

	log.Println("Running ...")
	time.Sleep(config.Duration())

	log.Println("Stopping peering  ...")
	for _, node := range nodeMap {
		node.Stop()
	}
	analysis.Stop()
	if config.VisEnabled() {
		stopServer()
	}
	log.Println("Stopping peering ... done")

	// Start finalize simulation result
	linkAnalysis := simulation.LinksToString(analysis.Links())
	err := simulation.WriteCSV(linkAnalysis, "linkAnalysis", []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	//	log.Println(linkAnalysis)

	convAnalysis := simulation.ConvergenceToString()
	err = simulation.WriteCSV(convAnalysis, "convAnalysis", []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	msgAnalysis := simulation.MessagesToString(nodeMap, analysis.Status())
	err = simulation.WriteCSV(msgAnalysis, "msgAnalysis", []string{"ID", "OUT", "ACC", "REJ", "IN", "DROP"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	err = simulation.WriteAdjlist(nodeMap, "adjlist")
	if err != nil {
		log.Fatalln("error writing adjlist:", err)
	}
}

func main() {
	config.Load()
	if err := logger.InitGlobalLogger(viper.GetViper()); err != nil {
		panic(err)
	}
	if config.VisEnabled() {
		startServer()
	}
	runSim()
}

func startServer() {
	srv = visualizer.NewServer()
	closing = make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.Run()
	}()
	log.Println("Server started; visit http://localhost:8844 to start")
	<-srv.Start
}

func stopServer() {
	close(closing)
	srv.Close()
	wg.Wait()
}

func statVisualizer() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-closing:
				return
			case <-ticker.C:
				visualizer.UpdateConvergence(simulation.RecordConv.GetConvergence())
				visualizer.UpdateAvgNeighbors(simulation.RecordConv.GetAvgNeighbors())
			}
		}
	}()
}
