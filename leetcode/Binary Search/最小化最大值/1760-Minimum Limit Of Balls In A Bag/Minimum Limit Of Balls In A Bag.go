package main

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	return sort.Search(1e9, func(i int) bool {
		i++
		cnt := 0
		for _, num := range nums {
			cnt += (num - 1) / i
		}
		return cnt <= maxOperations
	}) + 1
}
