package main

import (
	"math"
)

func maxProfit1(prices []int) int {
	n := len(prices)

	memo := make([][3][2]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, j, hold int) int {
		if j < 0 {
			return math.MinInt / 2
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return 0
		}
		if hold == 1 {
			memo[i][j][hold] = max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		} else {
			memo[i][j][hold] = max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])
		}
		return memo[i][j][hold]
	}
	return dfs(n-1, 2, 0)
}

func maxProfit2(prices []int) int {
	n := len(prices)

	f := make([][4][2]int, n+1)

	for i := range f {
		for j := range f[i] {
			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
		}
	}
	for j := 1; j <= 3; j++ {
		f[0][j][0] = 0
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= 3; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+prices[i])
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-prices[i])
		}
	}
	return f[n][3][0]
}

func maxProfit(prices []int) int {
	f := make([][2]int, 4)

	for j := 1; j <= 3; j++ {
		f[j][1] = math.MinInt / 2
	}
	for _, price := range prices {
		for j := 3; j > 0; j-- {
			f[j][1] = max(f[j][1], f[j-1][0]-price)
			f[j][0] = max(f[j][0], f[j][1]+price)
		}
	}
	return f[3][0]
}
