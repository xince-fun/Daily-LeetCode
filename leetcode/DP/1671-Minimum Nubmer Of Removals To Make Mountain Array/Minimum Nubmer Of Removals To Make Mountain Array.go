package main

import "sort"

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	suf := make([]int, n)
	g := []int{}
	// 从右往左
	for i := n - 1; i > 0; i-- {
		j := sort.SearchInts(g, nums[i])
		if j == len(g) {
			g = append(g, nums[i])
		} else {
			g[j] = nums[i]
		}
		suf[i] = j + 1 // 从 nums[i]开始的最长严格严格递减子序列的长度
	}

	mx := 0
	g = g[:0]
	for i, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) {
			g = append(g, x)
		} else {
			g[j] = x
		}
		pre := j + 1
		if pre >= 2 && suf[i] >= 2 {
			mx = max(mx, pre+suf[i]-1)
		}
	}
	return n - mx
}
