package main

import "math/bits"

func maximumRows1(matrix [][]int, numSelect int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	path := make([]bool, n)

	isValid := func(row []int, path []bool) bool {
		for i := 0; i < n; i++ {
			if row[i] == 1 && !path[i] {
				return false
			}
		}
		return true
	}

	var dfs func(int, int)
	dfs = func(i, chosen int) {
		if i == n {
			count := 0
			for j := 0; j < m; j++ {
				if isValid(matrix[j], path) {
					count++
				}
			}
			ans = max(ans, count)
			return
		}
		dfs(i+1, chosen)
		if chosen < numSelect {
			path[i] = true
			dfs(i+1, chosen+1)
			path[i] = false
		}
	}
	dfs(0, 0)
	return ans
}

func maximumRows(matrix [][]int, numSelect int) (ans int) {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i, row := range matrix {
		for j, x := range row {
			mask[i] |= x << j
		}
	}

	for subset := 0; subset < 1<<n; subset++ {
		if bits.OnesCount(uint(subset)) != numSelect {
			continue
		}
		coveredRows := 0
		for _, row := range mask {
			if row&subset == row {
				coveredRows++
			}
		}
		ans = max(ans, coveredRows)
	}
	return
}
