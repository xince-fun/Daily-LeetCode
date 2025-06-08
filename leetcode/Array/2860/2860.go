package main

import (
	"slices"
)

// [6,0,3,3,6,7,2,7]

// [0,2,3,3,6,6,7,7]

// [0,2,2,2,6,6,7,100]

func countWays(nums []int) (ans int) {
	n := len(nums)
	slices.Sort(nums)
	nums = append(nums, n+1)
	// 如果没有选人
	if 0 < nums[0] {
		ans++
	}
	for i := 0; i < n; i++ {
		// 选中的人数
		count := i + 1
		if count <= nums[i] {
			continue
		}
		// 选中的严格小于nums[i]
		if count < nums[i+1] {
			ans++
		}
	}
	return
}
