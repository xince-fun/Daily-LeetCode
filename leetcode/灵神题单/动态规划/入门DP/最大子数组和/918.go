package main

import "math"

func maxSubarraySum(nums []int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}

// 这种做法 O(n^2) 超时
func maxSubarraySumCircular1(nums []int) (ans int) {
	ans = math.MinInt
	n := len(nums)
	nums = append(nums, nums...)
	for i := 0; i < n; i++ {
		ans = max(ans, maxSubarraySum(nums[i:i+n]))
	}
	return
}

// 分类讨论一： 没有跨过边界，就是 最大子数组和
// 分类讨论二： 跨过边界，就是sum - 最小子数组和
// 特殊 当最小子数组和就是整个数组，此时最大子数组和就是空的，不符合，判断minS=sum,如果等，就只能是maxS

func maxSubarraySumCircular(nums []int) (ans int) {
	maxS := math.MinInt
	minS := 0 // 最小子数组和，可以为空
	maxf, minf, sum := 0, 0, 0
	for _, x := range nums {
		maxf = max(maxf, 0) + x
		maxS = max(maxS, maxf)
		minf = min(minf, 0) + x
		minS = min(minS, minf)
		sum += x
	}
	if sum == minS {
		return maxS
	}
	return max(maxS, sum-minS)
}
