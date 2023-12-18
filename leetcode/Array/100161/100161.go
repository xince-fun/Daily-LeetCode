package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	ans := [][]int{}

	sub := []int{}
	for i := 0; i < len(nums); {
		sub = append(sub, nums[i])

		if len(sub) == 3 {
			if sub[2]-sub[0] <= k {
				ans = append(ans, sub)
				sub = nil
			} else {
				return [][]int{}
			}
		}
		i++
	}
	if len(sub) > 0 {
		return [][]int{}
	}
	return ans
}
