package main

type Param struct {
	N            int64
	Known        float64
	vEnabled     bool
	SimDuration  int64
	DiscResStrat int64
}

func setParam(p *Param) {
	if p == nil {
		return
	}
	if p.N != 0 {
		N = int(p.N)
	}
	if p.Known != 0 {
		Known = p.Known
	}
	if p.SimDuration != 0 {
		SimDuration = int(p.SimDuration)
	}
	DiscResStrat = int(p.DiscResStrat)
	vEnabled = p.vEnabled
}
