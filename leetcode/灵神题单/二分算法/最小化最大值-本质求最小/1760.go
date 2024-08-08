package main

import "sort"

// 总共 9个球
// 每个袋里最多有 3个球, 也能最多4个球
// [3,3,3] [4,2,3] [5,1,3]
// [2,2,2] 不符合

// 随着 x 变大， 由false 变成 true

// 越靠近 right, operation的值越小 所以右边恒为true

func minimumSize(nums []int, maxOperations int) int {
	return sort.Search(1_000_000_001, func(x int) bool {
		cnt := 0
		for _, num := range nums {
			cnt += (num - 1) / (x + 1)
		}
		return cnt <= maxOperations
	}) + 1
}
