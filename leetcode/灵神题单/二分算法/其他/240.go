package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1 // 右上角开始
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++ // 这一行都小于target
		} else {
			j-- // 这一列都大于target
		}
	}
	return false
}
