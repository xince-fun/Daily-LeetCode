package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowFlag, colFlag := false, false
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowFlag = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowFlag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if colFlag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
}
