package main

import (
	"sort"
)

func splitArray1(nums []int, k int) int {
	sum, mx := 0, 0
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	left := max(mx, (sum-1)/k+1)
	right := sum
	ans := left + sort.Search(right-left, func(mx int) bool {
		mx += left
		cnt, s := 1, 0
		for _, x := range nums {
			if s+x <= mx {
				s += x
			} else {
				if cnt == k {
					return false
				}
				cnt += 1
				s = x
			}
		}
		return true
	})
	return ans
}

func splitArray(nums []int, k int) int {
	check := func(mx int) bool {
		cnt, s := 1, 0
		for _, x := range nums {
			if s+x <= mx {
				s += x
			} else {
				if cnt == k {
					return false
				}
				cnt += 1
				s = x
			}
		}
		return true
	}

	sum, mx := 0, 0
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	left := max(mx-1, (sum-1)/k)
	right := sum
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
