package main

import "sort"

func getTriggerTime(increase [][]int, requirements [][]int) (ans []int) {
	n := len(increase)
	// 前缀和
	c, r, h := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i, inc := range increase {
		c[i+1] = c[i] + inc[0]
		r[i+1] = r[i] + inc[1]
		h[i+1] = h[i] + inc[2]
	}

	for _, req := range requirements {
		// 找到>= c, r,h 的idx
		cIdx := sort.SearchInts(c, req[0])
		rIdx := sort.SearchInts(r, req[1])
		hIdx := sort.SearchInts(h, req[2])
		maxI := max(cIdx, rIdx, hIdx)
		if maxI == n+1 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, maxI)
		}
	}
	return
}
