package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
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

	fileMutex sync.Mutex
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

func appendToFile(f *os.File, s string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()
	f.WriteString(s)
}

func runSim(simCounter int) {

	f, err := os.OpenFile("data/peering-results.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	defer appendToFile(f, fmt.Sprintf("%s END SIMULATION #%d\n", time.Now().Format("2006/01/02 15:04:05.000000"), simCounter+1))

	adjFile, err := os.OpenFile("data/result_adjlist.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer adjFile.Close()

	appendToFile(f, fmt.Sprintf("%s BEGIN SIMULATION #%d\n", time.Now().Format("2006/01/02 15:04:05.000000"), simCounter+1))

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Run simulation with the following parameters:")
	config.PrintConfig()

	selection.SetParameters(selection.Parameters{
		InboundNeighborSize:    config.InboundNeighborhood(),
		OutboundNeighborSize:   config.OutboundNeighborhood(),
		SaltLifetime:           config.SaltLifetime(),
		OutboundUpdateInterval: 200 * time.Millisecond, // use exactly the same update time as previously
	})

	closure := events.NewClosure(func(ev *selection.PeeringEvent) {
		if ev.Status {
			log.Printf("Peering: %s<->%s\n", ev.Self.String(), ev.Peer.ID())
			appendToFile(f, fmt.Sprintf("%s Peering: %s<->%s\n", time.Now().Format("2006/01/02 15:04:05.000000"), ev.Self.String(), ev.Peer.ID()))
		}
	})
	selection.Events.OutgoingPeering.Attach(closure)
	defer selection.Events.OutgoingPeering.Detach(closure)

	closure1 := events.NewClosure(func(ev *selection.DroppedEvent) {
		log.Printf("Dropped: %s<->%s\n", ev.Self.String(), ev.DroppedID)
		appendToFile(f, fmt.Sprintf("%s Dropped: %s<->%s\n", time.Now().Format("2006/01/02 15:04:05.000000"), ev.Self.String(), ev.DroppedID))
	})
	selection.Events.Dropped.Attach(closure1)
	defer selection.Events.Dropped.Attach(closure1)

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
			config.Mana(),
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
	simulation.IdentityMana = simulation.NewZipfMana(identities, config.Zipf())
	for _, identity := range identities {
		if config.Mana() {
			appendToFile(f, fmt.Sprintf("%s ID - mana: %s - %d\n", time.Now().Format("2006/01/02 15:04:05.000000"), identity.ID(), simulation.IdentityMana[identity]))
		}
		if config.VisEnabled() {
			c := "0x666666"
			if config.Mana() {
				c = color(simulation.IdentityMana[identity], simulation.IdentityMana[identities[0]], simulation.IdentityMana[identities[len(identities)-1]])
			}
			visualizer.SetColor(identity.ID().String(), c)
		}
	}

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
	if config.VisEnabled() && config.Runs()-simCounter == 1 {
		stopServer()
	}
	log.Println("Stopping peering ... done")

	// Start finalize simulation result
	linkAnalysis := simulation.LinksToString(analysis.Links())
	err = simulation.WriteCSV(linkAnalysis, "linkAnalysis", []string{"X", "Y"})
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

	appendToFile(adjFile, fmt.Sprintf("### SIMULATION #%d\n", simCounter+1))
	err = simulation.WriteAdjlist(nodeMap, adjFile)
	if err != nil {
		log.Fatalln("error writing adjlist:", err)
	}

	// g := graph.New(identities)
	// for _, identity := range identities {
	// 	neighbors := nodeMap[identity.ID()].GetNeighbors()
	// 	for _, neighbor := range neighbors {
	// 		g.AddEdge(identity.ID().String(), neighbor.Identity.ID().String())
	// 	}
	// }
	// log.Println("Diameter: ", g.Diameter())
}

func main() {
	config.Load()
	if err := logger.InitGlobalLogger(viper.GetViper()); err != nil {
		panic(err)
	}
	if config.VisEnabled() {
		startServer()
	}

	// reset previous simulation results
	filenames := []string{"data/peering-results.txt", "data/result_adjlist.txt"}
	for _, filename := range filenames {
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		f.Close()
	}

	for i := 0; i < config.Runs(); i++ {
		runSim(i)
	}
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

func color(target, max, min uint64) string {
	base := "0x1100"
	//max-min : 255 = target-min : x
	v := (target - min) * 255 / (max - min)
	return base + fmt.Sprintf("%x", v)
}
