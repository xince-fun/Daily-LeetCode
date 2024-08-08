package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func closedIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	isClosed := true
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			isClosed = false
			return
		}
		if grid[x][y] == 2 || grid[x][y] == 1 {
			return
		}

		grid[x][y] = 2
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				isClosed = true
				dfs(i, j)
				if isClosed {
					ans++
				}
			}
		}
	}
	return
}
