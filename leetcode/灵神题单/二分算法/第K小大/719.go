package main

import "slices"

func smallestDistancePair(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)

	check := func(mid int) bool {
		sum := 0
		for i, j := 0, 1; i < n; i++ {
			for j < n && nums[j]-nums[i] <= mid {
				j++
			}
			sum += j - i - 1
		}

		return sum >= k
	}

	left, right := -1, nums[n-1]-nums[0]+1
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
