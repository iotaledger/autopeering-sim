package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/wollac/autopeering/peer"
	"github.com/wollac/autopeering/selection"
	"github.com/wollac/autopeering/simulation/visualizer"
)

var (
	allPeers           []*peer.Peer
	mgrMap             = make(map[peer.ID]*selection.Manager)
	idMap              = make(map[peer.ID]uint16)
	status             = NewStatusMap() // key: timestamp, value: Status
	neighborhoods      = make(map[peer.ID][]*peer.Peer)
	Links              = []Link{}
	DistInfo           = NewDistanceInfo()
	linkChan           = make(chan Event, 100)
	termTickerChan     = make(chan bool)
	termDistTickerChan = make(chan struct{})
	RecordConv         = NewConvergenceList()
	StartTime          time.Time
	wg                 sync.WaitGroup
	wgClose            sync.WaitGroup

	N            = 100
	vEnabled     = false
	SimDuration  = 300
	SaltLifetime = 300 * time.Second
	DropAllFlag  = false
	N_interval   = 1
	N_max        = 100
	numSim       = 30
	numStartSim  = 1 // 1st sim = 1, this value must  be smaller than numSim
)

func RunSim(loop int) {
	fileIndex := strconv.Itoa(loop)
	allPeers = make([]*peer.Peer, N)
	initialSalt := 0.
	//lambda := (float64(N) / SaltLifetime.Seconds()) * 10
	mgrMap = make(map[peer.ID]*selection.Manager)

	for i := range allPeers {
		peer := newPeer(fmt.Sprintf("%d", i), (time.Duration(initialSalt) * time.Second))
		allPeers[i] = peer.peer
		net := simNet{
			mgr:  mgrMap,
			loc:  peer.local,
			self: peer.peer,
			rand: peer.rand,
		}
		idMap[peer.local.ID()] = uint16(i)
		mgrMap[peer.local.ID()] = selection.NewManager(net, SaltLifetime, net.GetKnownPeers, peer.log, DropAllFlag)

		if vEnabled {
			visualizer.AddNode(peer.local.ID().String())
		}

		//initialSalt = initialSalt + (1 / lambda)				 // constant rate
		//initialSalt = initialSalt + rand.ExpFloat64()/lambda  // poisson process
		initialSalt = rand.Float64() * SaltLifetime.Seconds() // random
	}

	fmt.Println("start link analysis")
	runLinkAnalysis()
	runDistanceAnalysis()

	if vEnabled {
		statVisualizer()
	}

	StartTime = time.Now()
	for _, peer := range allPeers {
		mgrMap[peer.ID()].Start()
	}
	wgClose.Add(len(allPeers))

	fmt.Println("Sleep for ", time.Duration(SimDuration)*time.Second)
	time.Sleep(time.Duration(SimDuration) * time.Second)
	// Stop updating visualizer
	if vEnabled {
		termTickerChan <- true
	}
	// Stop simulation
	for _, p := range allPeers {
		p := p
		go func(p *peer.Peer) {
			mgrMap[p.ID()].Close()
			wgClose.Done()
		}(p)
	}
	log.Println("Closing...")
	wgClose.Wait()
	log.Println("Closing Done")
	linkChan <- Event{TERMINATE, 0, 0, 0}
	close(termDistTickerChan)

	// Wait until analysis goroutine stops
	wg.Wait()

	// Start finalize simulation result
	linkAnalysis := linksToString(LinkSurvival(Links))
	err := writeCSV(linkAnalysis, "linkAnalysis_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	//	log.Println(linkAnalysis)

	convAnalysis := convergenceToString(RecordConv.convergence)
	err = writeCSV(convAnalysis, "convAnalysis_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	//log.Println(RecordConv)

	msgAnalysis := messagesToString(status)
	err = writeCSV(msgAnalysis, "msgAnalysis_"+fileIndex, []string{"ID", "OUT", "ACC", "REJ", "IN", "DROP"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	distanceAnalysis, distanceAnalysisHisto := distanceToString()
	err = writeCSV(distanceAnalysis, "distanceAnalysis_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	distanceMedianAnalysis := distanceMedianToString(DistInfo)
	err = writeCSV(distanceMedianAnalysis, "distanceMedianAnalysis_"+fileIndex, []string{"X", "median"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	log.Println("Simulation Done")
	err = writeCSV(distanceAnalysisHisto, "distanceAnalysisHisto_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")

	distanceInboundAnalysis, distanceInboundAnalysisHisto := distanceInboundToString()
	err = writeCSV(distanceInboundAnalysis, "distanceInboundAnalysis_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")
	err = writeCSV(distanceInboundAnalysisHisto, "distanceInboundAnalysisHisto_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")

	distanceOutboundAnalysis, distanceOutboundAnalysisHisto := distanceOutboundToString()
	err = writeCSV(distanceOutboundAnalysis, "distanceOutboundAnalysis_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")
	err = writeCSV(distanceOutboundAnalysisHisto, "distanceOutboundAnalysisHisto_"+fileIndex, []string{"X", "Y"})
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}
	log.Println("Simulation Done")

	resetConvergence()

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
	for counter := numStartSim - 1; counter < numSim; counter++ {
		fmt.Println("... sim ", counter+1, "/", numSim)
		RunSim(counter)
	}
}

func runLinkAnalysis() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case newEvent := <-linkChan:
				switch newEvent.eType {
				case ESTABLISHED:
					Links = append(Links, NewLink(newEvent.x, newEvent.y, newEvent.timestamp.Milliseconds()))
					//log.Println("New Link", newEvent)
				case DROPPED:
					DropLink(newEvent.x, newEvent.y, newEvent.timestamp.Milliseconds(), Links)
					//log.Println("Link Dropped", newEvent)
				case TERMINATE:
					return
				}
			case <-ticker.C:
				updateConvergence(time.Since(StartTime))
			}
		}
	}()
}

func runDistanceAnalysis() {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			// time's up, get distance median
			case <-ticker.C:
				var dist []uint32
				// get all distances
				for _, p := range mgrMap {
					dist = append(dist, p.GetNeighborsDistance()...)
				}

				// sort the list
				sort.Slice(dist, func(i, j int) bool { return dist[i] < dist[j] })

				// get median
				medianIndex := len(dist) / 2
				var median uint32
				if len(dist)%2 == 1 {
					median = dist[medianIndex]
				} else {
					median = (dist[medianIndex] + dist[medianIndex+1]) / 2
				}

				// append median to the global list
				DistInfo.Append(median)

			case <-termDistTickerChan:
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
	for _, peer := range mgrMap {
		l := len(peer.GetNeighbors())
		if l == 8 {
			counter++
		}
		avgNeighbors += l
	}
	c := (float64(counter) / float64(N)) * 100
	avg := float64(avgNeighbors) / float64(N)
	RecordConv.Append(Convergence{time, c, avg})
}

func resetConvergence() {
	RecordConv = NewConvergenceList()

}
