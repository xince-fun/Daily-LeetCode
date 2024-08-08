package main

import "slices"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// 二分能到达的最远距离
	if ladders >= len(heights)-1 {
		return len(heights) - 1
	}
	left, right := ladders, len(heights)

	check := func(mid int, b, l int) bool {
		d := make([]int, mid)
		for i := 0; i < mid; i++ {
			d[i] = heights[i+1] - heights[i]
		}
		slices.Sort(d)
		for i := mid - 1; i >= 0; i-- {
			if d[i] < 0 {
				break
			}
			if l > 0 {
				l--
			} else if b >= d[i] {
				b -= d[i]
			} else {
				return false
			}
		}
		return true
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid, bricks, ladders) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
