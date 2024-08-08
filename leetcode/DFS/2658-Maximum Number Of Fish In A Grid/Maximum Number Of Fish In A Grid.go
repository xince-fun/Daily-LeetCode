package main

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || grid[i][j] == -1 {
			return
		}
		count += grid[i][j]
		grid[i][j] = -1
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				count = 0
				dfs(i, j)
				ans = max(ans, count)
			}
		}
	}
	return
}
