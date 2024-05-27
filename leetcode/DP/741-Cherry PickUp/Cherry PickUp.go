package main

import (
	"math"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)

	memo := make([][][]int, n*2-1)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, n)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(t, j, k int) int {
		if j < 0 || k < 0 || t < j || t < k || grid[t-j][j] < 0 || grid[t-k][k] < 0 {
			return math.MinInt
		}
		if t == 0 {
			return grid[0][0]
		}
		p := &memo[t][j][k]
		if *p != -1 {
			return *p
		}
		res := max(dfs(t-1, j, k), dfs(t-1, j, k-1), dfs(t-1, j-1, k), dfs(t-1, j-1, k-1)) + grid[t-j][j]
		if k != j {
			res += grid[t-k][k]
		}
		*p = res
		return res
	}
	return max(dfs(n*2-2, n-1, n-1), 0)
}

func cherryPickup1(grid [][]int) int {
	n := len(grid)
	f := make([][][]int, n*2-1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+1)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt
			}
		}
	}
	f[0][1][1] = grid[0][0]
	for t := 1; t < n*2-1; t++ {
		for j := max(t-n+1, 0); j <= min(t, n-1); j++ {
			if grid[t-j][j] < 0 {
				continue
			}
			for k := j; k <= min(t, n-1); k++ {
				if grid[t-k][k] < 0 {
					continue
				}
				f[t][j+1][k+1] = max(f[t-1][j+1][k+1], f[t-1][j+1][k], f[t-1][j][k+1], f[t-1][j][k]) + grid[t-j][j]
				if k != j {
					f[t][j+1][k+1] += grid[t-k][k]
				}
			}
		}
	}
	return max(f[n*2-2][n][n], 0)
}

func cherryPickup2(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[1][1] = grid[0][0]
	for t := 1; t < n*2-1; t++ {
		for j := min(t, n-1); j >= max(t-n+1, 0); j-- {
			for k := min(t, n-1); k >= j; k-- {
				if grid[t-j][j] < 0 || grid[t-k][k] < 0 {
					f[j+1][k+1] = math.MinInt
					continue
				}
				f[j+1][k+1] = max(max(f[j+1][k+1], f[j+1][k]), max(f[j][k+1], f[j][k])) + grid[t-j][j]
				if k != j {
					f[j+1][k+1] += grid[t-k][k]
				}
			}
		}
	}
	return max(f[n][n], 0)
}
