package kuji

import (
	"math/rand"
	"time"
)

type Util struct{}

/**
 * Spread all candidates into array
 * Array length is equals to sum total weight of candidates.
 * Values are Id for each candidates and spread in array equals to their weight value.
 */
func (u *Util) SpreadCandidates(candidates []KujiCandidate) []int64 {
	var values []int64
	for _, candidate := range candidates {

		for i := 0; i < int(candidate.Weight); i++ {
			values = append(values, candidate.Id)
		}
	}
	return values
}

/**
 * Shuffle elements using rand.Perm
 */
func (u *Util) Shuffle(target []int64) []int64 {
	n := len(target)
	dest := make([]int64, n)
	permutation := rand.Perm(n)
	for i, v := range permutation {
		dest[v] = target[i]
	}
	return dest
}

/**
 * Use Fisher–Yates shuffle
 * refs.
 * https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
 */
func (u *Util) Shuffle2(target []int64) []int64 {
	n := len(target)
	dest := make([]int64, n)
	copy(dest, target)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		dest[i], dest[j] = dest[j], dest[i]
	}
	return dest
}

/**
 * Get random number with min - max range.
 */
func (u *Util) RandomNumberFromRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
