package main

import "sort"

func minimumDifference(nums []int, k int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	ans = nums[k-1] - nums[0]
	for i := k; i < n; i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return
}
