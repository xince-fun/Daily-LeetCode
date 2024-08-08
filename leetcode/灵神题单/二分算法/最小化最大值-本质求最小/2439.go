package main

import (
	"slices"
	"sort"
)

func minimizeArrayValue(nums []int) int {
	mx := slices.Max(nums)

	return sort.Search(mx, func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
