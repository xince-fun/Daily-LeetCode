package main

import (
	"slices"
)

// 最大子数组和， 最小子数组和的绝对值(相反数)

func maxAbsoluteSum1(nums []int) (ans int) {
	n := len(nums)
	maxF, minF := make([]int, n), make([]int, n)
	maxF[0] = max(nums[0], 0)
	minF[0] = min(nums[0], 0)
	for i := 1; i < n; i++ {
		maxF[i] = max(maxF[i-1], 0) + nums[i]
		minF[i] = min(minF[i-1], 0) + nums[i]
	}
	return max(slices.Max(maxF), -slices.Min(minF))
}

func maxAbsoluteSum(nums []int) (ans int) {
	maxf, minf := 0, 0
	for _, x := range nums {
		maxf = max(maxf, 0) + x
		minf = min(minf, 0) + x
		ans = max(ans, max(maxf, -minf))
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
