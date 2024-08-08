package main

import "slices"

func smallestDivisor(nums []int, threshold int) int {
	maxn := slices.Max(nums)

	check := func(div int) bool {
		sum := 0
		for _, num := range nums {
			sum += (num + div - 1) / div
		}
		return sum <= threshold
	}

	left, right := 0, maxn+1
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
