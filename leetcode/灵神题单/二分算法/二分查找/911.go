package main

import "sort"

type pair struct {
	personId, timestamp int
}

type TopVotedCandidate struct {
	t []pair
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tArr := []pair{}
	maxVote, maxVoteId := -1, 0
	m := map[int]int{}
	for i, t := range times {
		id := persons[i]
		m[id]++
		if m[id] > maxVote {
			maxVote = m[id]
			maxVoteId = id
		}
		tArr = append(tArr, pair{maxVoteId, t})
	}
	return TopVotedCandidate{
		t: tArr,
	}

}

func (tc *TopVotedCandidate) Q(t int) int {
	// <= 时间t的idx
	idx := sort.Search(len(tc.t), func(i int) bool {
		return tc.t[i].timestamp >= t+1
	}) - 1
	return tc.t[idx].personId
}
