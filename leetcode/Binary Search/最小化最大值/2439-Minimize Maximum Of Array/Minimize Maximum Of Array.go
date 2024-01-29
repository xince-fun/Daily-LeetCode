package main

import "sort"

// 「最小化最大值」就是二分答案的代名词。

func minimizeArrayValue(nums []int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	return sort.Search(mx, func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
