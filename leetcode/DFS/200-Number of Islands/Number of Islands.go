package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0]) // m -> row n -> column
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, m, n, i, j)
			}
		}
	}
	return count
}

// 遍历过就变成 2
func dfs(grid [][]byte, m, n, x, y int) {
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
		return
	}

	grid[x][y] = '2'

	dfs(grid, m, n, x-1, y)
	dfs(grid, m, n, x+1, y)
	dfs(grid, m, n, x, y-1)
	dfs(grid, m, n, x, y+1)
}
