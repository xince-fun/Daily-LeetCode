package main

import "fmt"

// 选/不选

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
		// 不选
		dfs(i+1, chosen)
		// 选, path[i]标记为true
		if chosen < numSelect {
			path[i] = true
			dfs(i+1, chosen+1)
			path[i] = false
		}
	}
	dfs(0, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumRows(matrix [][]int, numSelect int) (ans int) {
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
		count := 0
		if chosen == numSelect {
			for j := 0; j < m; j++ {
				if isValid(matrix[j], path) {
					count++
				}
			}
			ans = max(ans, count)
		}
		if i == n {
			return
		}
		for j := i; j < n; j++ {
			path[j] = true
			dfs(j+1, chosen+1)
			path[j] = false
		}
	}
	dfs(0, 0)
	return
}

func main() {
	numSelect := 2
	matrix := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
	fmt.Println(maximumRows(matrix, numSelect))
}
