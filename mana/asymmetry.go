package main

import (
	"log"

	"github.com/iotaledger/hive.go/autopeering/mana"
)

func runAsymmetry(n, R int, ro, threshold, zipf float64) {
	IDs := newTestIdentities(n)
	testIdentityMana = newZipfMana(IDs, zipf)
	totalMana := mana.Total(manaF.Eval, IDs)

	// largest mana holder
	log.Println("largest mana holder ")

	target := IDs[0]

	set := mana.RankByFixedRange(manaF.Eval, target, IDs, R)
	log.Println("RankByFixedRange - len:", len(set))
	occurance := 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByFixedRange(manaF.Eval, identity, IDs, R))
	}
	log.Println("RankByFixedRange - asymmetry:", occurance)
	log.Printf("RankByFixedRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByVariableRange(manaF.Eval, target, IDs, R, ro)
	log.Println("RankByVariableRange - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByVariableRange(manaF.Eval, identity, IDs, R, ro))
	}
	log.Println("RankByVariableRange - asymmetry:", occurance)
	log.Printf("RankByVariableRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByThreshold(manaF.Eval, target, IDs, threshold)
	log.Println("RankByThreshold - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByThreshold(manaF.Eval, identity, IDs, threshold))
	}
	log.Println("RankByThreshold - asymmetry:", occurance)
	log.Printf("RankByThreshold - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	// smallest mana holder
	log.Println("smallest mana holder ")

	target = IDs[len(IDs)-1]

	set = mana.RankByFixedRange(manaF.Eval, target, IDs, R)
	log.Println("RankByFixedRange - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByFixedRange(manaF.Eval, identity, IDs, R))
	}
	log.Println("RankByFixedRange - asymmetry:", occurance)
	log.Printf("RankByFixedRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByVariableRange(manaF.Eval, target, IDs, R, ro)
	log.Println("RankByVariableRange - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByVariableRange(manaF.Eval, identity, IDs, R, ro))
	}
	log.Println("RankByVariableRange - asymmetry:", occurance)
	log.Printf("RankByVariableRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByThreshold(manaF.Eval, target, IDs, threshold)
	log.Println("RankByThreshold - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByThreshold(manaF.Eval, identity, IDs, threshold))
	}
	log.Println("RankByThreshold - asymmetry:", occurance)
	log.Printf("RankByThreshold - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	// middle mana holder
	log.Println("middle mana holder ")

	target = IDs[len(IDs)/2]

	set = mana.RankByFixedRange(manaF.Eval, target, IDs, R)
	log.Println("RankByFixedRange - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByFixedRange(manaF.Eval, identity, IDs, R))
	}
	log.Println("RankByFixedRange - asymmetry:", occurance)
	log.Printf("RankByFixedRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByVariableRange(manaF.Eval, target, IDs, R, ro)
	log.Println("RankByVariableRange - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByVariableRange(manaF.Eval, identity, IDs, R, ro))
	}
	log.Println("RankByVariableRange - asymmetry:", occurance)
	log.Printf("RankByVariableRange - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)

	set = mana.RankByThreshold(manaF.Eval, target, IDs, threshold)
	log.Println("RankByThreshold - len:", len(set))
	occurance = 0
	for _, identity := range set {
		occurance += asymmetryCheck(target, mana.RankByThreshold(manaF.Eval, identity, IDs, threshold))
	}
	log.Println("RankByThreshold - asymmetry:", occurance)
	log.Printf("RankByThreshold - mana: %.2f %% \n", float64(mana.Total(manaF.Eval, set))/float64(totalMana-manaF.Eval(target))*100)
}
