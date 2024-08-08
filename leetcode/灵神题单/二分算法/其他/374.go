package main

import "sort"

func guess(num int) int {
	return 0
}

// 右边的都是大于等于的

func guessNumber1(n int) int {
	left, right := 0, n+1
	for left+1 < right {
		mid := left + (right-left)/2
		r := guess(mid)
		if r <= 0 {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func guessNumber(n int) int {
	return sort.Search(n, func(i int) bool {
		return guess(i) <= 0
	})
}
