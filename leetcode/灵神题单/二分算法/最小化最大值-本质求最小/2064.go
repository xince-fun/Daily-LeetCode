package main

import (
	"slices"
	"sort"
)

// [11, 6]
// 每个店最多 3 个商品 最多4个
// [2,3,3,3,3,3] [2,4,2,3,3,3]

// 右边恒为true

func minimizedMaximum1(n int, quantities []int) int {
	// 二分的左边界从1开始
	return sort.Search(1e5, func(max int) bool {
		cnt := 0
		for _, q := range quantities {
			cnt += (q + max) / (max + 1)
		}
		return cnt <= n
	}) + 1
}

func minimizedMaximum(n int, quantities []int) int {
	left := 0
	right := slices.Max(quantities)

	for left+1 < right {
		mid := left + (right-left)/2
		sum := len(quantities)
		for _, p := range quantities {
			sum += (p - 1) / mid
		}
		if sum <= n {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
