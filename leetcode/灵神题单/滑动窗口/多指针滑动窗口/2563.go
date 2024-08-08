package main

import (
	"slices"
	"sort"
)

func countFairPairs1(nums []int, lower int, upper int) (ans int64) {
	slices.Sort(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return
}

func countFairPairs(nums []int, lower int, upper int) (ans int64) {
	slices.Sort(nums)
	left, right := len(nums), len(nums)
	for j, x := range nums {
		for right > 0 && nums[right-1] > upper-x {
			right--
		}
		for left > 0 && nums[left-1] >= lower-x {
			left--
		}
		ans += int64(min(right, j) - min(left, j))
	}
	return
}
