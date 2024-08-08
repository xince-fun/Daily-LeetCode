package main

import "math"

func minOperations(nums []int, x int) int {
	total := 0
	ans := math.MinInt
	for _, num := range nums {
		total += num
	}
	left, sum := 0, 0
	target := total - x
	for right, num := range nums {
		sum += num
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			ans = max(ans, right-left+1)
		}
	}
	if ans == math.MinInt {
		return -1
	}
	return len(nums) - ans
}
