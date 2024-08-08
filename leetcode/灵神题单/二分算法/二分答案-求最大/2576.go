package main

import (
	"slices"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	return 2 * sort.Search(n/2, func(k int) bool {
		k++
		for i, x := range nums[:k] {
			if x*2 > nums[n-k+i] {
				return true
			}
		}
		return false
	})
}
