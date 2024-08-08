package main

import "sort"

func shipWithinDays1(weights []int, days int) int {
	return sort.Search(2500*1e4, func(i int) bool {
		cnt := 1
		sum := 0
		for _, w := range weights {
			if w > i {
				return false
			}
			sum += w
			if sum > i {
				cnt++
				sum = w
			}
		}
		return cnt <= days
	})
}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 25000000

	check := func(ww int) bool {
		cnt := 1
		sum := 0
		for _, w := range weights {
			if w > ww {
				return false
			}
			sum += w
			if sum > ww {
				cnt++
				sum = w
			}
		}
		return cnt <= days
	}

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
