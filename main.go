package main

import (
	"github.com/iotaledger/hive.go/configuration"
	"github.com/iotaledger/hive.go/events"
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
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/logger"
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

func runSim(configuration *configuration.Configuration) {
	log.Println("Run simulation with the following parameters:")
	configuration.Print()

	selection.SetParameters(selection.Parameters{
		ArRowLifetime:          config.ArrowLifetime(configuration),
		OutboundUpdateInterval: 200 * time.Millisecond, // use exactly the same update time as previously
	})

	//lambda := (float64(N) / ArrowLifetime.Seconds()) * 10
	initialSalt := 0.

	log.Println("Creating peers ...")
	netw := transport.NewNetwork()
	nodeMap = make(map[identity.ID]simulation.Node, config.NumberNodes(configuration))
	for i := 0; i < config.NumberNodes(configuration); i++ {
		node := simulation.NewNode(transport.PeerID(i), time.Duration(initialSalt)*time.Second, netw, config.DropOnUpdate(configuration), discover)
		nodeMap[node.ID()] = node

		if config.VisEnabled(configuration) {
			visualizer.AddNode(node.ID().String())
		}

		// initialSalt = initialSalt + (1 / lambda)				 // constant rate
		// initialSalt = initialSalt + rand.ExpFloat64()/lambda  // poisson process
		initialSalt = rand.Float64() * config.ArrowLifetime(configuration).Seconds() // random
		inClosure := events.NewClosure(func(ev *selection.PeeringEvent) {
			if ev.Status {
				log.Printf("Accepted peering: %s<->%s (chan: %d)\n", node.Peer().ID(), ev.Peer.ID(), ev.Channel)
			} else {
				log.Printf("Rejected peering: %s<->%s (chan: %d)\n", node.Peer().ID(), ev.Peer.ID(), ev.Channel)
			}
		})

		node.Prot.Events().IncomingPeering.Attach(inClosure)

		dropClosure := events.NewClosure(func(ev *selection.DroppedEvent) {
			log.Printf("Dropping: %s<->%s\n", node.Peer().ID(), ev.DroppedID)
		})
		node.Prot.Events().Dropped.Attach(dropClosure)
	}

	log.Println("Creating peers ... done")

	analysis := simulation.NewLinkAnalysis(nodeMap, configuration)

	if config.VisEnabled(configuration) {
		statVisualizer()
	}

	log.Println("Starting peering ...")
	analysis.Start()
	for _, node := range nodeMap {
		node.Start()
	}
	log.Println("Starting peering ... done")

	log.Println("Running ...")
	time.Sleep(config.Duration(configuration))

	log.Println("Stopping peering  ...")
	for _, node := range nodeMap {
		node.Stop()
	}
	analysis.Stop()
	if config.VisEnabled(configuration) {
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

	msgAnalysis := simulation.MessagesToString(nodeMap, analysis.Status(), configuration)
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
	config := configuration.New()
	_ = config.LoadFile("config.json")

	if err := logger.InitGlobalLogger(config); err != nil {
		panic(err)
	}
	if config.Bool("VisualEnabled") {
		startServer()
	}
	runSim(config)
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
