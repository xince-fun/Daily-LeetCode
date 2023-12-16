package main

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	st := []int{}
	k := math.MinInt32
	for i := n - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
			x := st[len(st)-1]
			st = st[:len(st)-1]
			k = max(k, nums[x])
		}
		st = append(st, i)
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
