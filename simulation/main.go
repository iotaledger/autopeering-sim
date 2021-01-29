package main

import (
	"fmt"
	"log"
	//"math/rand"
	"sync"
	"time"

	"github.com/iotaledger/autopeering-sim/peer"
	"github.com/iotaledger/autopeering-sim/selection"
	"github.com/iotaledger/autopeering-sim/server"
	"github.com/iotaledger/autopeering-sim/simulation/visualizer"
	"github.com/iotaledger/autopeering-sim/transport"
	"github.com/iotaledger/hive.go/events"
)

var (
	allPeers       []*peer.Peer
	peerMap        = make(map[peer.ID]*peer.Peer)
	protocolMap    = make(map[peer.ID]*selection.Protocol)
	idMap          = make(map[peer.ID]uint16)
	status         = NewStatusMap() // key: timestamp, value: Status
	convMsg        []int
	neighborhoods  = make(map[peer.ID][]*peer.Peer)
	Links          = []Link{}
	termTickerChan = make(chan bool)
	incomingChan   = make(chan *selection.PeeringEvent, 10)
	outgoingChan   = make(chan *selection.PeeringEvent, 10)
	dropChan       = make(chan *selection.DroppedEvent, 10)
	closing        = make(chan struct{})
	saltTermChan   = make(chan bool)
	RecordConv     = NewConvergenceList()
	msgPerTList    = make(map[uint16][]int)
	StartTime      time.Time
	wg             sync.WaitGroup

	N            = 100
	vEnabled     = false
	SimDuration  = 300
	SaltLifetime = 300 * time.Second
	DropAllFlag  = false
)

// dummyDiscovery is a dummy implementation of DiscoveryProtocol.
type dummyDiscovery struct{}

// This simulation focuses on peer selection only.
// Thus, every peer is verified in order to skip the discovery ping-pong process.
func (d dummyDiscovery) IsVerified(peer.ID, string) bool                    { return true }
func (d dummyDiscovery) EnsureVerified(p *peer.Peer)                        {}
func (d dummyDiscovery) GetVerifiedPeer(id peer.ID, addr string) *peer.Peer { return peerMap[id] }
func (d dummyDiscovery) GetVerifiedPeers() []*peer.Peer                     { return allPeers }

func RunSim() {
	allPeers = make([]*peer.Peer, N)

	network := transport.NewNetwork()
	serverMap := make(map[peer.ID]*server.Server, N)
	disc := dummyDiscovery{}

	// subscribe to the events
	selection.Events.IncomingPeering.Attach(events.NewClosure(func(e *selection.PeeringEvent) { incomingChan <- e }))
	selection.Events.OutgoingPeering.Attach(events.NewClosure(func(e *selection.PeeringEvent) { outgoingChan <- e }))
	selection.Events.Dropped.Attach(events.NewClosure(func(e *selection.DroppedEvent) { dropChan <- e }))

	convMsg = make([]int, N)
	initialSalt := 0.
	//lambda := (float64(N) / SaltLifetime.Seconds()) * 10

	log.Println("Creating peers...")
	for i := range allPeers {
		name := fmt.Sprintf("%d", i)
		network.AddTransport(name)

		peer := newPeer(name, (time.Duration(initialSalt) * time.Second))
		allPeers[i] = peer.peer

		id := peer.local.ID()
		idMap[id] = uint16(i)
		peerMap[id] = peer.peer

		cfg := selection.Config{Log: peer.log,
			Param: &selection.Parameters{
				SaltLifetime:          SaltLifetime,
				DropNeighborsOnUpdate: DropAllFlag,
			},
		}
		protocol := selection.New(peer.local, disc, cfg)
		serverMap[id] = server.Listen(peer.local, network.GetTransport(name), peer.log, protocol)

		protocolMap[id] = protocol

		if vEnabled {
			visualizer.AddNode(id.String())
		}

		// initialSalt = initialSalt + (1 / lambda)				 // constant rate
		// initialSalt = initialSalt + rand.ExpFloat64()/lambda  // poisson process
		//initialSalt = rand.Float64() * SaltLifetime.Seconds() // random
	}

	fmt.Println("start link analysis")
	runLinkAnalysis()
	runMsgAnalysis()

	if vEnabled {
		statVisualizer()
	}

	StartTime = time.Now()
	for _, peer := range allPeers {
		srv := serverMap[peer.ID()]
		protocolMap[peer.ID()].Start(srv)
	}

	/* Get stable phase info
	    // Remove the collected data of the first T in order to collect data of stable phase,
	    // i.e., peers already have neighbors, and started to update their neighbors.
	    // TODO: make this configurable, and use T instead of 30.
		time.Sleep(time.Duration(30) * time.Second)
		for i := range allPeers {
			status.ClearStatusMap(uint16(i))
			RecordConv.ClearConvergence()
		}

		time.Sleep(time.Duration(SimDuration-30) * time.Second)
	*/

	// sleep here and wait the peers finding neighbors
	time.Sleep(time.Duration(SimDuration) * time.Second)

	// Stop updating visualizer
	if vEnabled {
		termTickerChan <- true
	}

	// Stop all peers
	log.Println("Closing...")
	for _, peer := range allPeers {
		protocolMap[peer.ID()].Close()
	}
	log.Println("Closing Done")

	// stop runLinkAnalysis and finalize convMsg data
	close(closing)
	// stop runMsgAnalysis and finalize msgPerT data
	saltTermChan <- true

	// Wait until analysis goroutine stops
	wg.Wait()

	// Conclude analysis results and write into csv files
	writeAnalysisResults()

	log.Println("Simulation Done")
}

func main() {
	p := parseInput("input.txt")
	setParam(p)

	var s *visualizer.Server
	if vEnabled {
		s = visualizer.NewServer()
		go s.Run()
		<-s.Start
	}
	fmt.Println("start sim")
	RunSim()
}

func runLinkAnalysis() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {

			// handle incoming peering requests
			case req := <-incomingChan:
				from := idMap[req.Peer.ID()]
				to := idMap[req.Self]
				status.Append(from, to, INCOMING)

			// handle outgoing peering requests
			case req := <-outgoingChan:
				from := idMap[req.Self]
				to := idMap[req.Peer.ID()]
				status.Append(from, to, OUTBOUND)

				// accepted/rejected is only recorded for outgoing requests
				if req.Status {
					status.Append(from, to, ACCEPTED)
					Links = append(Links, NewLink(from, to, time.Since(StartTime).Milliseconds()))
					if vEnabled {
						visualizer.AddLink(req.Self.String(), req.Peer.ID().String())
					}
				} else {
					status.Append(from, to, REJECTED)
				}

			// handle dropped peers incoming and outgoing
			case req := <-dropChan:
				from := idMap[req.Self]
				to := idMap[req.DroppedID]
				status.Append(from, to, DROPPED)
				DropLink(from, to, time.Since(StartTime).Milliseconds(), Links)
				if vEnabled {
					visualizer.RemoveLink(req.Self.String(), req.DroppedID.String())
				}

			case <-ticker.C:
				updateConvergence(time.Since(StartTime))

			case <-closing:
				for _, p := range allPeers {
					id := idMap[p.ID()]
					convMsg[id] = status.MsgLen(id)
				}
				return
			}
		}
	}()
}

func runMsgAnalysis() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case p := <-selection.ExpiredSaltChan:
				pID := idMap[p]
				msgPerTList[pID] = append(msgPerTList[pID], status.MsgLen(pID))
				status.ClearStatusMap(pID)
			case <-saltTermChan:
				// if some node's salt is not expired
				for i := range allPeers {
					msgPerTList[uint16(i)] = append(msgPerTList[uint16(i)], status.MsgLen(uint16(i)))
				}
				return
			}
		}
	}()
}

func statVisualizer() {
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		defer wg.Done()
		for {
			select {
			case <-termTickerChan:
				return
			case <-ticker.C:
				visualizer.UpdateConvergence(RecordConv.GetConvergence())
				visualizer.UpdateAvgNeighbors(RecordConv.GetAvgNeighbors())
			}
		}
	}()
}

func updateConvergence(time time.Duration) {
	counter := 0
	avgNeighbors := 0
	for _, prot := range protocolMap {
		l := len(prot.GetNeighbors())
		if l == 8 {
			counter++
		}
		avgNeighbors += l
	}
	c := (float64(counter) / float64(N)) * 100
	avg := float64(avgNeighbors) / float64(N)
	RecordConv.Append(Convergence{time, c, avg})
}

func writeAnalysisResults() {
	// Start finalize simulation result
	linkAnalysis := linksToString(LinkSurvival(Links))
	err := writeCSV(linkAnalysis, "linkAnalysis", []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	convAnalysis := convergenceToString(RecordConv.convergence)
	err = writeCSV(convAnalysis, "convAnalysis", []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	msgAnalysis := messagesToString(status)
	err = writeCSV(msgAnalysis, "msgAnalysis", []string{"ID", "OUT", "ACC", "REJ", "IN", "DROP"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	// calculate avg messages to converge
	convMsgAnalysis, convAvgMsg := convMsgToString(convMsg)
	err = writeCSV(convMsgAnalysis, "convMsgAnalysis", []string{"MSG"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	fmt.Println("avg message to converge ", convAvgMsg)

	// calculate avg messages per T
	msgPerTAnalysis, msgPerTAvg := msgPerTToString()
	err = writeCSV(msgPerTAnalysis, "msgPerTAnalysis", []string{"MSG"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	fmt.Println("avg message per T ", msgPerTAvg)
}
