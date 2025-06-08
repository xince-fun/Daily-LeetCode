package main

import (
	"math"
	"slices"
)

// dfs(i, j) = min(dfs(i-1,j),dfs(i-1,j-1)) + grid[i][j]

func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		f[i][0] = math.MaxInt
	}
	for j := 0; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	f[0][0] = 0
	for i, row := range triangle {
		for j := 0; j < i; j++ {
			f[i+1][j+1] = min(f[i][j], f[i][j+1]) + row[j]
		}
		f[i+1][i+1] = f[i][i] + row[i]
	}
	return slices.Min(f[n])
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt
	}
	f[0] = 0
	for i, row := range triangle {
		if i != 0 {
			f[0] = math.MaxInt
		}
		f[i+1] = f[i] + row[i]
		for j := i - 1; j >= 0; j-- {
			f[j+1] = min(f[j], f[j+1]) + row[j]
		}
	}
	return slices.Min(f)
}
