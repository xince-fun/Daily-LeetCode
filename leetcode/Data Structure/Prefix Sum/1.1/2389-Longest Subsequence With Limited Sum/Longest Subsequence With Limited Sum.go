package main

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for i, query := range queries {
		queries[i] = sort.SearchInts(nums, query+1)
	}
	return queries
}
