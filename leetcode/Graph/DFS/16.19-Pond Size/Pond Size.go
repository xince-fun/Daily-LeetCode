package main

import "sort"

func pondSizes(land [][]int) (ans []int) {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || land[x][y] != 0 {
			return 0
		}
		land[x][y] = 1
		ans := 1

		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				ans += dfs(i, j)
			}
		}
		return ans
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 0 {
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	return
}
