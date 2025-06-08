package main

import (
	"math"
	"slices"
)

// dfs(i,j) = grid[i][j] + min(range dfs(i-1，j))

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+2)
		f[i][0], f[i][n+1] = math.MaxInt, math.MaxInt
	}
	for i, row := range grid {
		for j, x := range row {
			res := math.MaxInt
			for k := 0; k < n; k++ {
				if k != j {
					res = min(res, f[i][k])
				}
			}
			f[i+1][j+1] = res + x
		}
	}
	return slices.Min(f[n])
}
