package main

import "sort"

func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	n := len(nums)

	left, right := 0, n-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}
