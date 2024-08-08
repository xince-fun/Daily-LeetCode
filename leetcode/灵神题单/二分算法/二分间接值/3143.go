package main

import (
	"math/bits"
	"sort"
)

func maxPointsInsideSquare(points [][]int, s string) (ans int) {
	// 找到
	sort.Search(1_000_000_001, func(size int) bool {
		vis := 0
		for i, p := range points {
			if abs(p[0]) <= size && abs(p[1]) <= size {
				c := s[i] - 'a'
				if vis>>c&1 > 0 {
					return true
				}
				vis |= 1 << c
			}
		}
		ans = bits.OnesCount(uint(vis))
		return false
	})
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
