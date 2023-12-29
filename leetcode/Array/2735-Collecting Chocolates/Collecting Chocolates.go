package main

import "slices"

// (i + 1) mod n )
// 这个操作就是 数组左移一次

// [20, 1, 15]  => [20, 1, 15, 20,]
// 计算次数 => 20   1    15
//			  ^    ^    ^
//		      2    0   1
//
// 比如现在是  [14, 20, 30, 1, 15,  16]
//             ^   ^   ^  ^   ^    ^
//			   3   4   5  0   1    2

func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int64, n)
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums {
		for j := i; j < n+i; j++ {
			mn = min(mn, nums[j%n])
			s[j-i] += int64(mn)
		}
	}
	return slices.Min(s)
}

func minCost1(nums []int, x int) int64 {

	return 0
}
