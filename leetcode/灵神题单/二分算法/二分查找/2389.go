package main

import (
	"slices"
	"sort"
)

// <= 转换为 (>x)-1 即 (>= x+1)-1

func answerQueries(nums []int, queries []int) []int {
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for i, x := range queries {
		// 这里的 x + 1的意思是 找到 >= x+1的位置， 外面其实要-1，但是长度加1， 即为这样
		queries[i] = sort.SearchInts(nums, x+1)
	}
	return queries
}
