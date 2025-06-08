package main

import (
	"math"
	"slices"
)

// dfs(i,j) = min(dfs(i-1,j-1),dfs(i-1,j),dfs(i-1,j+1)) + matrix[i][j]

// 解决方法：在左边和右边各加一行

func minFallingPathSum1(matrix [][]int) int {
	n := len(matrix)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+2)
		f[i][0] = math.MaxInt
		f[i][n+1] = math.MaxInt
	}
	for i, row := range matrix {
		for j, x := range row {
			f[i+1][j+1] = min(f[i][j], f[i][j+1], f[i][j+2]) + x
		}
	}
	return slices.Min(f[n])
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	f := make([]int, n+2)
	for _, row := range matrix {
		pre := f[0]
		for j, x := range row {
			pre, f[j+1] = f[j+1], min(pre, f[j+1], f[j+2])+x
		}
	}
	return slices.Min(f)
}
