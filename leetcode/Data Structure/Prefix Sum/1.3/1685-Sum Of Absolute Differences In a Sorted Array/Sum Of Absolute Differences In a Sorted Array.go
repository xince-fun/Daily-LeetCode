package main

import "sort"

func getSumAbsoluteDifferences1(nums []int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int, n)
	for i := range ans {
		j := sort.SearchInts(nums, nums[i])
		left := j*nums[i] - sum[j]
		right := sum[n] - sum[j] - nums[i]*(n-j)
		ans[i] = left + right
	}
	return ans
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int, n)
	for i := range ans {
		left := (i+1)*nums[i] - sum[i+1]
		right := sum[n] - sum[i] - nums[i]*(n-i)
		ans[i] = left + right
	}
	return ans
}
