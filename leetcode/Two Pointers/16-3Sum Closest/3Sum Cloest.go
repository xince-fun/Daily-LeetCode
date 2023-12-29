package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	diff := math.MaxInt
	for i := 0; i < n-2; i++ {
		x := nums[i]
		j, k := i+1, n-1
		for j < k {
			sum := x + nums[j] + nums[k]
			if sum > target {
				k--
			} else if sum < target {
				j++
			} else {
				return sum
			}
			if abs(sum-target) < diff {
				diff = abs(sum - target)
				ans = sum
			}
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
