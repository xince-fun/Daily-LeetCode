package main

import (
	"math"
	"slices"
)

// dfs(i,j)表示到达坐标i,j 所需的最小代价
// dfs(i,j) = min(dfs(i-1, \range j) + moveCost[x][j])

//dfs(i,j)=grid[i][j]+ k=0minn−1dfs(i+1,k)+moveCost[grid[i][j]][k]

func minPathCost1(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[m-1] = grid[m-1]
	for i := m - 2; i >= 0; i-- {
		for j, g := range grid[i] {
			f[i][j] = math.MaxInt
			for k, c := range moveCost[g] {
				f[i][j] = min(f[i][j], f[i+1][k]+c)
			}
			f[i][j] += g
		}
	}
	return slices.Min(f[0])
}

func minPathCost(grid [][]int, moveCost [][]int) int {
	for i := len(grid) - 2; i >= 0; i-- {
		for j, g := range grid[i] {
			res := math.MaxInt
			for k, c := range moveCost[g] {
				res = min(res, grid[i+1][k]+c)
			}
			grid[i][j] += res
		}
	}
	return slices.Min(grid[0])
}
