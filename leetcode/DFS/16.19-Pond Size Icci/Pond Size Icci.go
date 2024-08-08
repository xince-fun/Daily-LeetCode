package main

import (
	"slices"
)

func pondSizes(land [][]int) (ans []int) {
	m, n := len(land), len(land[0])

	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || (land[i][j] > 0 && land[i][j] <= 1000) || land[i][j] == -1 {
			return
		}
		land[i][j] = -1
		count++
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j+1)
		dfs(i-1, j-1)
		dfs(i+1, j-1)
		dfs(i-1, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				count = 0
				dfs(i, j)
				ans = append(ans, count)
			}
		}
	}
	slices.Sort(ans)
	return
}
