package main

import "math"

//

// dfs(i, 0)

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
			memo[i][hold] = max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
		} else {
			memo[i][hold] = max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
		}
		return memo[i][hold]
	}

	return dfs(n-1, 0)
}

func maxProfit2(prices []int) int {
	n := len(prices)
	f := make([][2]int, n+2)
	f[0][0] = 0
	f[1][1] = math.MinInt

	for i := 0; i < n; i++ {
		f[i+2][1] = max(f[i+1][1], f[i][0]-prices[i])
		f[i+2][0] = max(f[i+1][0], f[i+1][1]+prices[i])
	}
	return f[n+1][0]
}

func maxProfit(prices []int) int {
	pre0, f0, f1 := 0, 0, math.MinInt
	for _, price := range prices {
		new_f0 := max(f0, f1+price)
		f1 = max(f1, pre0-price)
		pre0 = f0
		f0 = new_f0

	}
	return f0
}
