package main

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)
	ans := []int{}

	for len(nums) > 0 {
		x, y := nums[0], nums[1]
		ans = append(ans, y, x)
		nums = nums[2:]
	}
	return ans
}
