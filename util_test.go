package kuji

import (
	"testing"
)

var util *Util

func TestUtil_SpreadCandidates(t *testing.T) {
	result := util.SpreadCandidates([]KujiCandidate{
		KujiCandidate{
			Id:     1,
			Weight: 100,
		},
		KujiCandidate{
			Id:     2,
			Weight: 200,
		},
		KujiCandidate{
			Id:     3,
			Weight: 700,
		},
	})

	if len(result) != (100 + 200 + 700) {
		t.Errorf("Result array length not enough or too few (len: %d)", len(result))
	}
	id1Cnt := 0
	id2Cnt := 0
	id3Cnt := 0
	for _, v := range result {
		switch v {
		case 1:
			id1Cnt++
		case 2:
			id2Cnt++
		case 3:
			id3Cnt++
		}
	}
	if id1Cnt != 100 {
		t.Errorf("Candidates number mismatched of Id=1 (cnt: %d)", id1Cnt)
	}
	if id2Cnt != 200 {
		t.Errorf("Candidates number mismatched of Id=2 (cnt: %d)", id2Cnt)
	}
	if id3Cnt != 700 {
		t.Errorf("Candidates number mismatched of Id=3 (cnt: %d)", id3Cnt)
	}
}

func TestUtil_RandomNumberFromRange(t *testing.T) {
	result := util.RandomNumberFromRange(13, 81)

	if !(result >= 13 && result <= 81) {
		t.Errorf("Random number out of range (result: %d)", result)
	}
}

func TestUtil_Shuffle(t *testing.T) {
	before := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	after := util.Shuffle(before)
	matchedCnt := 0
	for i, v := range after {
		if v == before[i] {
			matchedCnt++
		}
	}
	if matchedCnt == len(before) {
		t.Errorf("Array not shuffled (matched: %d / len: %d)", matchedCnt, len(before))
	}
}

func TestUtil_Shuffle2(t *testing.T) {
	before := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	after := util.Shuffle2(before)
	matchedCnt := 0
	for i, v := range after {
		if v == before[i] {
			matchedCnt++
		}
	}
	if matchedCnt == len(before) {
		t.Errorf("Array not shuffled (matched: %d / len: %d)", matchedCnt, len(before))
	}
}
