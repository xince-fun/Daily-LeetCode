package main

import (
	"slices"
	"sort"
)

// 设 f(d) 表示 磁力为d时能放置的球
// 磁力越大，能放置的位置越少，具有单调性，能够二分

// d_0为答案
// f(d_0) >= m 且 f(d_0+1) < k

func maxDistance(position []int, m int) int {
	slices.Sort(position)
	right := (position[len(position)-1] - position[0]) / (m - 1)
	return sort.Search(right, func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		cnt, pre := 1, position[0]
		for _, p := range position[1:] {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt < m
	})
}
