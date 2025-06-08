package main

import "math"

// dfs(-1,j) dfs(i, -1) = 0

func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = grid[0][0]
	for i, row := range grid {
		for j, x := range row {
			if i == 0 && j > 0 {
				f[i][j] = f[i][j-1] + x
			} else if j == 0 && i > 0 {
				f[i][j] = f[i-1][j] + x
			} else if i > 0 && j > 0 {
				f[i][j] = min(f[i-1][j], f[i][j-1]) + x
			}
		}
	}
	return f[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		f[i][0] = math.MaxInt
	}
	for j := 0; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	f[0][1] = 0
	for i, row := range grid {
		for j, x := range row {
			f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + x
		}
	}
	return f[m][n]
}

func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := [2][]int{make([]int, n+1), make([]int, n+1)}
	f[0][0], f[1][0] = math.MaxInt, math.MaxInt
	for j := 0; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	f[0][1] = 0
	for i, row := range grid {
		for j, x := range row {
			f[(i+1)%2][j+1] = min(f[i%2][j+1], f[(i+1)%2][j]) + x
		}
	}
	return f[m%2][n]
}

func minPathSum4(grid [][]int) int {
	n := len(grid[0])
	f := make([]int, n+1)
	for j := 1; j <= n; j++ {
		f[j] = math.MaxInt
	}
	for i, row := range grid {
		if i != 0 {
			f[0] = math.MaxInt
		}
		for j, x := range row {
			f[j+1] = min(f[j+1], f[j]) + x
		}
	}
	return f[n]
}

// 原地修改

// f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]

func minPathSum(grid [][]int) int {
	n := len(grid[0])
	f := grid[0]
	for j := 1; j < n; j++ {
		f[j] += f[j-1]
	}
	for _, row := range grid[1:] {
		f[0] += row[0]
		for j := 1; j < n; j++ {
			f[j] = min(f[j], f[j-1]) + row[j]
		}
	}
	return f[n-1]
}
