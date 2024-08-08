package main

import "slices"

// 超内存
func findKthNumber1(m int, n int, k int) int {
	arr := make([]int, m*n)
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[index] = (i + 1) * (j + 1)
			index++
		}
	}
	slices.Sort(arr)
	return arr[k-1]
}

func findKthNumber(m int, n int, k int) int {

	check := func(mid int) bool {
		i, j := m-1, 0
		sum := 0
		for i >= 0 && j < n {
			num := (i + 1) * (j + 1)
			if num <= mid {
				sum += i + 1
				j++
			} else {
				i--
			}
		}
		return sum >= k
	}

	left, right := -1, (m+1)*(n+1)+1
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
