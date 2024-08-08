package main

import "sort"

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion1(n int) int {
	left, right := -1, n+1
	for left+1 < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func firstBadVersion2(n int) int {
	return sort.Search(n+1, func(i int) bool {
		return isBadVersion(i)
	})
}

func firstBadVersion(n int) int {
	return sort.Search(n+1, func(i int) bool {
		return isBadVersion(i)
	})
}
