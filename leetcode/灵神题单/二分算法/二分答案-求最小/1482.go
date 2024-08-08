package main

import "slices"

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	left := 0
	right := slices.Max(bloomDay)
	for left+1 < right {
		mid := left + (right-left)/2

		idx := 0
		cnt := 0
		for idx < n {
			if bloomDay[idx] > mid {
				idx++
				continue
			}
			could := true
			for i := 0; i < k; i++ {
				if idx >= n {
					could = false
					break
				}
				if bloomDay[idx] > mid {
					could = false
					break
				}
				idx++
			}
			if could {
				cnt++
			}
		}
		if cnt >= m {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
