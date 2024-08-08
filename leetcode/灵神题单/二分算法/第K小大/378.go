package main

import (
	"slices"
)

func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix)
	index := 0
	arr := make([]int, n*n)
	for _, row := range matrix {
		for _, x := range row {
			arr[index] = x
			index++
		}
	}
	slices.Sort(arr)
	return arr[k-1]
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	check := func(mid int) bool {
		i, j := n-1, 0
		sum := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				sum += i + 1
				j++
			} else {
				i--
			}
		}
		return sum >= k
	}

	left, right := matrix[0][0]-1, matrix[n-1][n-1]+1
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
