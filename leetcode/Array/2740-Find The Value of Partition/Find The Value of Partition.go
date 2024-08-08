package main

import (
	"math"
	"slices"
)

func findValueOfPartition(nums []int) (ans int) {
	slices.Sort(nums)
	ans = math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return
}
