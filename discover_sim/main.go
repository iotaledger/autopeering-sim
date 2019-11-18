package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/iotaledger/autopeering-sim/discover"
	"github.com/iotaledger/autopeering-sim/peer"
	"github.com/iotaledger/autopeering-sim/server"
	"github.com/iotaledger/autopeering-sim/transport"
)

var (
	allPeers    []simPeer
	protocolMap = make(map[peer.ID]*discover.Protocol)
	idMap       = make(map[peer.ID]uint16)
	closing     = make(chan struct{})
	RecordKnown = NewConvergenceList()
	SentMsg     = make(map[uint16][]byte)
	StartTime   time.Time
	wg          sync.WaitGroup

	N           = 100
	vEnabled    = false
	SimDuration = 300
	Known       = 0.01
)

func RunSim() {
	allPeers = make([]simPeer, N)

	network := transport.NewNetwork()
	serverMap := make(map[peer.ID]*server.Server, N)
	initialSalt := 0.
	numEntry := int(float64(N) * Known)

	log.Println("Create peers...")
	for i := range allPeers {
		name := fmt.Sprintf("%d", i)
		network.AddTransport(name)
		allPeers[i] = newPeer(name, (time.Duration(initialSalt) * time.Second))
	}

	log.Println("Initiate peers...")
	for i, p := range allPeers {
		name := fmt.Sprintf("%d", i)

		id := p.local.ID()
		idMap[id] = uint16(i)

		var cfg discover.Config
		var boot []*peer.Peer
		for j := 0; j < numEntry; j++ {
			if i != j {
				boot = append(boot, allPeers[j].peer)
			}
		}

		cfg = discover.Config{Log: p.log,
			MasterPeers: boot,
		}
		protocol := discover.New(p.local, cfg)
		serverMap[id] = server.Listen(p.local, network.GetTransport(name), p.log, protocol)

		protocolMap[id] = protocol
	}

	Analysis()

	StartTime = time.Now()
	for _, p := range allPeers {
		srv := serverMap[p.peer.ID()]
		protocolMap[p.peer.ID()].Start(srv)
	}
	// entry nodes first connect to each other
	/*
	    fmt.Println(time.Now())
	    for i := 0; i < numEntry; i++ {
			srv := serverMap[allPeers[i].peer.ID()]
			protocolMap[allPeers[i].peer.ID()].Start(srv)
		}
		time.Sleep(time.Duration(40) * time.Second)

	    fmt.Println(time.Now())
	    for i := numEntry; i < N; i++ {
			srv := serverMap[allPeers[i].peer.ID()]
			protocolMap[allPeers[i].peer.ID()].Start(srv)
		}
	*/

	time.Sleep(time.Duration(SimDuration) * time.Second)

	// Stop simulation
	log.Println("Closing...")
	for _, p := range allPeers {
		serverMap[p.peer.ID()].Close()
		protocolMap[p.peer.ID()].Close()
		p.db.Close()
	}
	close(closing)
	log.Println("Closing Done")

	// Wait until analysis goroutine stops
	wg.Wait()

	knownPeerAnalysis := knownPeersToString(RecordKnown.convergence)
	err := writeCSV(knownPeerAnalysis, "knownPeerAnalysis", []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	sentMsgAnalysis := sentMsgToString(SentMsg)
	err = writeCSV(sentMsgAnalysis, "sentMsgAnalysis", []string{"ID", "MSG"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	sentMsgPdfAnalysis := sentMsgPdfToString(SentMsg)
	err = writeCSV(sentMsgPdfAnalysis, "sentMsgPdfAnalysis", []string{"MSG", "%"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")
}

func main() {
	p := parseInput("input.txt")
	setParam(p)
	print(p.SimDuration)

	fmt.Println("start sim")
	RunSim()
}

func Analysis() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				updateNumKnownPeers(time.Since(StartTime))
			case m := <-discover.UpdateMsg:
				p := idMap[m.GetPeer()]
				SentMsg[p] = append(SentMsg[p], m.GetType())
			case <-closing:
				return
			}
		}
	}()
}

func updateNumKnownPeers(time time.Duration) {
	counter := 0
	avgKnown := 0
	for _, prot := range protocolMap {
		l := len(prot.GetKnownPeers())
		//l := len(prot.GetVerifiedPeers())
		if l == (N - 1) {
			counter++
		}
		avgKnown += l
	}
	avg := float64(avgKnown) / float64(N)
	RecordKnown.Append(Convergence{time, counter, avg})
}
