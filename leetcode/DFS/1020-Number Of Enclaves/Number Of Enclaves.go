package main

func numEnclaves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return
}
