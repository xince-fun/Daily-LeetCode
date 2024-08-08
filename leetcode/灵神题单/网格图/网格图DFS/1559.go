package main

var dirs = []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	var dfs func(byte, int, int, int, int) bool

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	dfs = func(c byte, x, y, lx, ly int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != c {
			return false
		}
		if vis[x][y] {
			return true
		}
		vis[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx == lx && ny == ly {
				continue
			}
			if dfs(c, nx, ny, x, y) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] && dfs(grid[i][j], i, j, -1, -1) {
				return true
			}
		}
	}
	return false
}
