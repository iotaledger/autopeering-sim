package simulation

import (
	"fmt"
	"sync"
	"time"

	"github.com/iotaledger/goshimmer/packages/autopeering/peer"
)

const (
	ACCEPTED = 'A'
	REJECTED = 'R'
	DROPPED  = 'D'
	OUTBOUND = 'O'
	INCOMING = 'I'
)

var RecordConv = NewConvergenceList()

type Status struct {
	timestamp int64
	opType    byte
	toNode    peer.ID
}

type StatusSum struct {
	outbound int
	accepted int
	incoming int
	rejected int
	dropped  int
}

type Link struct {
	x, y         peer.ID
	tEstablished int64
	tDropped     int64
}

type Convergence struct {
	timestamp    time.Duration
	counter      float64
	avgNeighbors float64
}

type ConvergenceList struct {
	sync.Mutex
	convergence []Convergence
}

type StatusMap struct {
	sync.Mutex
	status map[peer.ID][]Status
}

func NewConvergenceList() *ConvergenceList {
	return &ConvergenceList{
		convergence: []Convergence{},
	}
}

func (c *ConvergenceList) Append(t Convergence) {
	c.Lock()
	defer c.Unlock()
	c.convergence = append(c.convergence, t)
}

func (c *ConvergenceList) GetConvergence() float64 {
	c.Lock()
	defer c.Unlock()
	cLen := len(c.convergence)
	if cLen > 0 {
		return c.convergence[cLen-1].counter
	}
	return 0
}

func (c *ConvergenceList) GetAvgNeighbors() float64 {
	c.Lock()
	defer c.Unlock()
	cLen := len(c.convergence)
	if cLen > 0 {
		return c.convergence[cLen-1].avgNeighbors
	}
	return 0
}

func NewStatusMap() *StatusMap {
	return &StatusMap{
		status: make(map[peer.ID][]Status),
	}
}

func (s *StatusMap) Append(from, to peer.ID, op byte) {
	s.Lock()
	defer s.Unlock()
	st := Status{
		timestamp: time.Now().Unix(),
		opType:    op,
		toNode:    to,
	}
	s.status[from] = append(s.status[from], st)
}

func (s *StatusMap) GetSummary(id peer.ID) (cnt StatusSum) {
	s.Lock()
	defer s.Unlock()
	for _, t := range s.status[id] {
		switch t.opType {
		case ACCEPTED:
			cnt.accepted++
		case REJECTED:
			cnt.rejected++
		case DROPPED:
			cnt.dropped++
		case OUTBOUND:
			cnt.outbound++
		case INCOMING:
			cnt.incoming++
		}
	}
	return cnt
}

func NewLink(x, y peer.ID, timestamp int64) Link {
	return Link{
		x:            x,
		y:            y,
		tEstablished: timestamp,
	}
}

func DropLink(x, y peer.ID, timestamp int64, list []Link) bool {
	for i := len(list) - 1; i >= 0; i-- {
		if (list[i].x == x && list[i].y == y) ||
			(list[i].x == y && list[i].y == x) {
			if list[i].tDropped == 0 {
				list[i].tDropped = timestamp
				return true
			}
			return false
		}
	}
	return false
}

func (l Link) String() string {
	result := ""
	result += fmt.Sprintf("\n%d--%d\t", l.x, l.y)
	if l.tDropped == 0 {
		return result
	}
	result += fmt.Sprintf("%d", l.tDropped-l.tEstablished)
	return result
}

func linkSurvival(links []Link) map[int64]int {
	result := make(map[int64]int)
	for _, l := range links {
		if l.tDropped != 0 {
			result[(l.tDropped-l.tEstablished)/1000]++
		}
	}
	return result
}
