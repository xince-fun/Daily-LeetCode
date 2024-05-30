package main

import "math"

func minMoves(nums []int, k int) int {
	p := []int{}
	for i, v := range nums {
		if v != 0 {
			p = append(p, i-len(p))
		}
	}
	m := len(p)
	s := make([]int, m+1)
	for i, v := range p {
		s[i+1] = s[i] + v
	}
	ans := math.MaxInt
	for i, v := range s[:m-k+1] {
		ans = min(ans, v+s[i+k]-s[i+k/2]*2-p[i+k/2]*(k%2))
	}
	return ans
}
