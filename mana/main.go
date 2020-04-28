package main

import (
	"fmt"
	"log"
	"math"

	"github.com/iotaledger/hive.go/autopeering/mana"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/hive.go/identity"
)

func newTestIdentity(name string) *identity.Identity {
	key := ed25519.PublicKey{}
	copy(key[:], name)
	return identity.New(key)
}

func newTestIdentities(n int) (identities []*identity.Identity) {
	identities = make([]*identity.Identity, n)
	for i := 0; i < n; i++ {
		identities[i] = newTestIdentity(fmt.Sprintf("%d", i))
	}
	return
}

var testIdentityMana map[*identity.Identity]uint64

func newTestMana(identities []*identity.Identity) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(identities))
	for i, p := range identities {
		m[p] = uint64(i)
	}
	return m
}

func newZipfMana(identities []*identity.Identity, zipf float64) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(identities))
	scalingFactor := math.Pow(10, 10)
	for i, p := range identities {
		m[p] = uint64(math.Pow(float64(i+1), -zipf) * scalingFactor)
		//log.Println(m[p])
	}
	return m
}

type sameManaFunc mana.Func

var sameMana sameManaFunc

func (f sameManaFunc) Eval(identity *identity.Identity) uint64 {
	return 1
}

type manaFunc mana.Func

var manaF manaFunc

func (f manaFunc) Eval(identity *identity.Identity) uint64 {
	return testIdentityMana[identity]
}

func stringID(identities []*identity.Identity) (output []string) {
	for _, item := range identities {
		output = append(output, fmt.Sprintf(item.ID().String()))
	}
	return
}

func asymmetryCheck(target *identity.Identity, identities []*identity.Identity) int {
	for _, identity := range identities {
		if identity == target {
			return 0
		}
	}
	return 1
}

func TestZipfMana() {
	IDs := newTestIdentities(100)
	testIdentityMana = newZipfMana(IDs, 0.82)
	totalMana := mana.Total(manaF.Eval, IDs)
	log.Println("Total Mana:", totalMana)

	for _, id := range IDs {
		fmt.Printf("%s - %.4f\n", id.ID().String(), float64(manaF.Eval(id))/float64(totalMana))
	}
}

func runAsymmetry() {
	IDs := newTestIdentities(1000)
	testIdentityMana = newZipfMana(IDs, 0.82)
	totalMana := mana.Total(manaF.Eval, IDs)

	R := 10
	ro := 1.2
	threshold := 1. / 3.

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

func main() {
	runAsymmetry()
}
