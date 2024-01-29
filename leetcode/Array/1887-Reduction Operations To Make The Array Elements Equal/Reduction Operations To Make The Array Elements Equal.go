package main

import "sort"

func reductionOperations(nums []int) (ans int) {
	sort.Ints(nums)
	for i, k, n := 0, 0, len(nums); i < n; k++ {
		start := i
		for ; i < n && nums[i] == nums[start]; i++ {
		}
		ans += (i - start) * k
	}
	return
}
