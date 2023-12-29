package main

import "math"

// dfs(i, j, 0) 表示到第i天结束时完成至多j笔交易，未持有股票的最大利润
// dfs(i, j, 1) 表示到第i天结束时完成至多j笔交易，持有股票的最大利润

// dfs(i, j, 0) = max(dfs(i-1,j,0), dfs(i-1,j,1)+prices[i])
// dfs(i, j, 1) = max(dfs(i-1,j,1), dfs(i-1,j-1,0)-prices[i])

// 递归边界
// dfs(·,-1,·) = -∞ // 任何情况下，j都不能为负
// dfs(-1,j,0) = 0 // 第0天开始未持有股票
// dfs(-1,,1) = -∞ // 第0天不可能持有股票

// dfs(n-1,k,0)

func maxProfit1(k int, prices []int) int {
	n := len(prices)

	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k+1)
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
		if memo[i][j][hold] != -1 {
			return memo[i][j][hold]
		}
		if hold == 1 {
			memo[i][j][hold] = max(dfs(i-1, j, 1), dfs(i-1, j-1, 0)-prices[i])
		} else {
			memo[i][j][hold] = max(dfs(i-1, j, 0), dfs(i-1, j, 1)+prices[i])
		}
		return memo[i][j][hold]
	}
	return dfs(n-1, k, 0)
}

// f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1] + prices[i])
// f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0] - prices[i])

// f[·][0][·] = -∞
// f[0][j][0] = 0 j >= 1
// f[0][j][1] = -∞ j >= 1
// f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+prices[i])
// f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-prices[i])

// 恰好 f[0][1][0] = 0 其余为 -∞

// 至少 f[i][-1][·] 等价于 f[i][0][·]
// 所以每个f[i]前面不需要插入状态
// [至少0次] 等价于可以无限次交易
// f[i][0][·] 就是无限次交易下的最大利润

// f[0][0][0] = 0, 其余 = -∞
// f[i+1][0][0] = max(f[i][0][0], f[i][0][1] + prices[i])
// f[i+1][0][1] = max(f[i][0][1], f[i][0][0] - prices[i])

func maxProfit2(k int, prices []int) int {
	n := len(prices)

	f := make([][][2]int, n+1)
	for i := range f {
		f[i] = make([][2]int, k+2)
		for j := range f[i] {
			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
		}
	}
	for j := 1; j <= k+1; j++ {
		f[0][j][0] = 0
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= k+1; j++ {
			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+prices[i])
			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-prices[i])
		}
	}
	return f[n][k+1][0]
}

func maxProfit(k int, prices []int) int {
	f := make([][2]int, k+2)

	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2

	for _, price := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+price)
			f[j][1] = max(f[j][1], f[j-1][0]-price)
		}
	}
	return f[k+1][0]
}
