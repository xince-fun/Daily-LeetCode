package main

// dfs(i,j) = dfs(i-1,j), dfs(i,j-1) * grid[i][j]
// dfs(i,j) = max(dfs(i-1,j) * grid[i][j], dfs(i,j-1) * grid[i][j])

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = 1
		}
	}
	for i, row := range grid {
		for j, x := range row {
			f[i+1][j+1] = max(f[i][j+1]*x, f[i][j-1]*x)
		}
	}
	return f[m][n]
}
