package main

import "slices"

func countSubarrays(nums []int, k int) int64 {
	maxN := slices.Max(nums)
	left, cnt := 0, 0
	ans := 0
	for _, x := range nums {
		if x == maxN {
			cnt++
		}
		for cnt >= k {
			if nums[left] == maxN {
				cnt--
			}
			left++
		}
		ans += left
	}
	return int64(ans)
}
