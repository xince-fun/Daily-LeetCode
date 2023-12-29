package main

import (
	"math"
)

// 子问题？ 到第i天结束时，持有/未持有股票的最大利润
// 下一个子问题？ 到i-1天结束时，持有/未持有股票的最大利润

// dfs(i, 0) 到i天结束时未持有股票
// dfs(i, 1) 表示到i天结束时持有股票

// 由于第i-1天的结束就是第i天的开始
// dfs(i-1, ·)表示到第i天开始时的最大利润

// dfs(i, 0) = max(dfs(i-1,0), dfs(i-1,1) + prices[i])
// dfs(i, 1) = max(dfs(i-1,1), dfs(i-1,0) - prices[i])

// dfs(-1, 0) = 0
// dfs(-1, 1) = -∞

func maxProfit1(prices []int) int {
	n := len(prices)

	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}

	var dfs func(int, int) int
	dfs = func(i int, hold int) int {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return 0
		}
		if memo[i][hold] != -1 {
			return memo[i][hold]
		}
		if hold == 1 {
			memo[i][hold] = max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		} else {
			memo[i][hold] = max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
		}
		return memo[i][hold]
	}

	return dfs(n-1, 0)
}

func maxProfit2(prices []int) int {
	n := len(prices)
	f := make([][2]int, n+1)
	f[0][0] = 0
	f[0][1] = math.MinInt

	for i := 0; i < n; i++ {
		f[i+1][1] = max(f[i][1], f[i][0]-prices[i])
		f[i+1][0] = max(f[i][0], f[i][1]+prices[i])
	}
	return f[n][0]
}

func maxProfit(prices []int) int {
	f0, f1 := 0, math.MinInt
	for _, price := range prices {
		new_f0 := max(f0, f1+price)
		f1 = max(f1, f0-price)
		f0 = new_f0
	}
	return f0
}
