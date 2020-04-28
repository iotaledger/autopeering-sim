package main

import (
	"log"

	"github.com/iotaledger/hive.go/autopeering/mana"
)

func runDistribution(n, R int, ro, threshold, zipf float64) {
	IDs := newTestIdentities(n)
	testIdentityMana = newZipfMana(IDs, zipf)

	fixedDistr := make(map[int]int)
	variableDistr := make(map[int]int)
	thresholdDistr := make(map[int]int)

	for _, id := range IDs {
		set := mana.RankByFixedRange(manaF.Eval, id, IDs, R)
		fixedDistr[len(set)]++
		makePlot(fixedDistr, "RankByFixedRange", "RankByFixedRange.png")

		set = mana.RankByVariableRange(manaF.Eval, id, IDs, R, ro)
		variableDistr[len(set)]++
		makePlot(variableDistr, "RankByVariableRange", "RankByVariableRange.png")

		set = mana.RankByThreshold(manaF.Eval, id, IDs, threshold)
		thresholdDistr[len(set)]++
		makePlot(thresholdDistr, "RankByThreshold", "RankByThreshold.png")
	}

	log.Println(fixedDistr)
	log.Println(variableDistr)
	log.Println(thresholdDistr)

}
