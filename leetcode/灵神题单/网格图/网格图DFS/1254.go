package main

func closedIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	isClosed := true
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			isClosed = false
			return
		}
		if grid[i][j] == 2 || grid[i][j] == 1 {
			return
		}

		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
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
