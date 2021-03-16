package simulation

import (
	"math"

	"github.com/iotaledger/hive.go/autopeering/mana"
	"github.com/iotaledger/hive.go/identity"
)

type manaFunc mana.Func

var manaF manaFunc

var IdentityMana map[*identity.Identity]uint64

func (f manaFunc) Eval(identity *identity.Identity) uint64 {
	return IdentityMana[identity]
}

func NewZipfMana(identities []*identity.Identity, zipf float64) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(identities))
	scalingFactor := math.Pow(10, 10)
	for i, p := range identities {
		m[p] = uint64(math.Pow(float64(i+1), -zipf) * scalingFactor)
		//log.Println(m[p])
	}
	return m
}

func NewPollenMana(identities []*identity.Identity, cm ConsensusMana) (m map[*identity.Identity]uint64) {
	m = make(map[*identity.Identity]uint64, len(cm.Consensus))

	for i, manaEntry := range cm.Consensus {
		m[identities[i]] = uint64(manaEntry.Mana)
	}
	return m
}
