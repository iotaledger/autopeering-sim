package simulation

import (
	"sync"
	"time"

	"github.com/iotaledger/autopeering-sim/simulation/config"
	"github.com/iotaledger/autopeering-sim/simulation/visualizer"
	"github.com/iotaledger/hive.go/autopeering/selection"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/identity"
)

type linkAnalysis struct {
	nodeMap map[identity.ID]Node

	startTime    time.Time
	incomingChan chan *selection.PeeringEvent
	outgoingChan chan *selection.PeeringEvent
	dropChan     chan *selection.DroppedEvent
	status       *StatusMap
	links        []Link

	closing chan struct{}
	wg      sync.WaitGroup
}

func NewLinkAnalysis(nodeMap map[identity.ID]Node) *linkAnalysis {
	return &linkAnalysis{
		nodeMap:      nodeMap,
		incomingChan: make(chan *selection.PeeringEvent, 10),
		outgoingChan: make(chan *selection.PeeringEvent, 10),
		dropChan:     make(chan *selection.DroppedEvent, 10),
		status:       NewStatusMap(),
		closing:      make(chan struct{}),
	}
}

func (la *linkAnalysis) Start() {
	la.startTime = time.Now()

	la.wg.Add(1)
	go la.loop()
}

func (la *linkAnalysis) Stop() {
	close(la.closing)
	la.wg.Wait()
}

func (la *linkAnalysis) Links() []Link {
	return la.links
}

func (la *linkAnalysis) Status() *StatusMap {
	return la.status
}

func (la *linkAnalysis) loop() {
	defer la.wg.Done()

	// start listening to the events
	defer la.subscribe()()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {

		// handle incoming peering requests
		case req := <-la.incomingChan:
			la.status.Append(req.Peer.ID(), req.Self, INCOMING)

		// handle outgoing peering requests
		case req := <-la.outgoingChan:
			from := req.Self
			to := req.Peer.ID()
			la.status.Append(from, to, OUTBOUND)

			// accepted/rejected is only recorded for outgoing requests
			if req.Status {
				la.status.Append(from, to, ACCEPTED)
				la.links = append(la.links, NewLink(from, to, time.Since(la.startTime).Milliseconds()))
				if config.VisEnabled() {
					visualizer.AddLink(req.Self.String(), req.Peer.ID().String())
				}
			} else {
				la.status.Append(from, to, REJECTED)
			}

		// handle dropped peers incoming and outgoing
		case req := <-la.dropChan:
			from := req.Self
			to := req.DroppedID
			la.status.Append(from, to, DROPPED)
			DropLink(from, to, time.Since(la.startTime).Milliseconds(), la.links)
			if config.VisEnabled() {
				visualizer.RemoveLink(req.Self.String(), req.DroppedID.String())
			}

		case <-ticker.C:
			la.updateConvergence(time.Since(la.startTime))

		case <-la.closing:
			return
		}
	}
}

// subscribe subscribes to the selection events.
func (la *linkAnalysis) subscribe() func() {
	incomingClosure := events.NewClosure(func(e *selection.PeeringEvent) { la.incomingChan <- e })
	outgoingClosure := events.NewClosure(func(e *selection.PeeringEvent) { la.outgoingChan <- e })
	dropClosure := events.NewClosure(func(e *selection.DroppedEvent) { la.dropChan <- e })

	selection.Events.IncomingPeering.Attach(incomingClosure)
	selection.Events.OutgoingPeering.Attach(outgoingClosure)
	selection.Events.Dropped.Attach(dropClosure)

	return func() {
		selection.Events.IncomingPeering.Detach(incomingClosure)
		selection.Events.OutgoingPeering.Detach(outgoingClosure)
		selection.Events.Dropped.Detach(dropClosure)
	}
}

func (la *linkAnalysis) updateConvergence(time time.Duration) {
	counter := 0
	avgNeighbors := 0
	for _, node := range la.nodeMap {
		l := len(node.GetNeighbors())
		if l == selection.DefaultOutboundNeighborSize+selection.DefaultInboundNeighborSize {
			counter++
		}
		avgNeighbors += l
	}
	c := (float64(counter) / float64(config.NumberNodes())) * 100
	avg := float64(avgNeighbors) / float64(config.NumberNodes())
	RecordConv.Append(Convergence{time, c, avg})
}
