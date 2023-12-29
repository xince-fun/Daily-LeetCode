package main

import (
	"sort"
)

func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)

	// nums[j] + nums[k] > x
	for i := n - 1; i > 1; i-- {
		x := nums[i]
		j, k := 0, i-1
		for j < k {
			if nums[j]+nums[k] > x {
				ans += k - j
				k--
			} else {
				j++
			}
		}
	}
	return
}
