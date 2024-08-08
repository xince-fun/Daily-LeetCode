package main

func maxAreaOfIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)

	count := 0
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		count++
		ans = max(ans, count)
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count = 0
				dfs(i, j)
			}
		}
	}
	return
}
