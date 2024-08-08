package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) (ans int) {
	sort.Ints(nums)
	ans = math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}
