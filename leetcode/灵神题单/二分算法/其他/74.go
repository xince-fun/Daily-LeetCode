package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	left, right := -1, m*n
	for left+1 < right {
		mid := left + (right-left)/2
		num := matrix[mid/n][mid%n]
		if target == num {
			return true
		}
		if num < target {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}
