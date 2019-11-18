package main

import (
	//"fmt"
	"sync"
	"time"
)

const (
	PING    = 'I'
	PONG    = 'O'
	DISCREQ = 'Q'
	DISCRES = 'S'
)

type Convergence struct {
	timestamp time.Duration
	counter   int
	avgKnown  float64
}

type ConvergenceList struct {
	sync.Mutex
	convergence []Convergence
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
