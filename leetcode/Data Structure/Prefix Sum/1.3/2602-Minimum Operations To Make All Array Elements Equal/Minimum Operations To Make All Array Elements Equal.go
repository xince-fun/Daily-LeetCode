package main

import "sort"

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - sum[j]               // 蓝色面积
		right := sum[n] - sum[j] - q*(n-j) // 绿色面积
		ans[i] = int64(left + right)
	}
	return ans
}
