package kuji

import (
	"math/rand"
	"time"
)

type Util struct {}

func (u *Util) SpreadCandidates(c []KujiCandidate) []int64 {
	var values []int64
	for _, candidate := range c {

		for i := 0; i < int(candidate.Weight); i++ {
			values = append(values, candidate.Id)
		}
	}
	return values
}

func (u *Util) RandomNumberFromRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max - min) + min
}
