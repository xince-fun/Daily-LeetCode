package main

import (
	"cmp"
	"slices"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {
	type pair struct{ x, y int }
	pairs := make([]pair, len(aliceValues))
	for i, x := range aliceValues {
		pairs[i] = pair{x, bobValues[i]}
	}
	slices.SortFunc(pairs, func(p, q pair) int { return q.x + q.y - p.x - p.y })
	diff := 0
	for i, p := range pairs {
		if i%2 == 0 {
			diff += p.x
		} else {
			diff -= p.y
		}
	}
	return cmp.Compare(diff, 0)
}
