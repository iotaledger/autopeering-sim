package main

import "time"

type Param struct {
	N           int64
	vEnabled    bool
	SimDuration int64
	T           int64
	dropAll     bool
    N_interval  int64
    N_max       int64
}

func setParam(p *Param) {
	if p == nil {
		return
	}
	if p.N != 0 {
		N = int(p.N)
        N_max = N
	}
	if p.T != 0 {
		SaltLifetime = time.Duration(p.T) * time.Second
	}
	if p.SimDuration != 0 {
		SimDuration = int(p.SimDuration)
	}
    if p.N_interval != 0 && p.N_max != 0 {
        N_interval = int(p.N_interval)
        N_max = int(p.N_max)
    }
	vEnabled = p.vEnabled
	DropAllFlag = p.dropAll
}
