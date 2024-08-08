package main

import (
	"slices"
	"sort"
)

func minEatingSpeed1(piles []int, h int) int {
	return 1 + sort.Search(slices.Max(piles)-1, func(i int) bool {
		cnt := 0
		i++
		for _, p := range piles {
			cnt += (p + i - 1) / i
		}
		return cnt <= h
	})
}

func minEatingSpeed2(piles []int, h int) int {
	left := 0
	right := slices.Max(piles)

	check := func(m int) bool {
		cnt := 0
		for _, p := range piles {
			cnt += (p + m - 1) / m
		}
		return cnt <= h
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

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	left := 0
	right := slices.Max(piles)

	for left+1 < right {
		mid := left + (right-left)/2
		sum := n
		for _, p := range piles {
			sum += (p - 1) / mid
		}
		if sum <= h {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
