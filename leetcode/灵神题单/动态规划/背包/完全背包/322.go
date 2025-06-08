package main

import "math"

/*
capacity: 背包容量
w[i]: 第 i 个物品的体积
v[i]: 第 i 个物品的价值
每种物品可以无限次重复选
返回： 所选物品体积和不超过 capacity 前提下，所能得到的最大价值和
*/

func unboundedKnapSack(capacity int, w, v []int) int {
	n := len(w)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity+1)
		for j := 0; j <= capacity; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			return 0
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		if c < w[i] {
			memo[i][c] = dfs(i-1, c)
		}
		memo[i][c] = max(dfs(i-1, c), dfs(i, c-w[i])+v[i])
		return memo[i][c]
	}
	return dfs(n-1, capacity)
}

func coinChange1(coins []int, amount int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		if c < coins[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = min(dfs(i-1, c), dfs(i, c-coins[i])+1)
		}
		return memo[i][c]
	}

	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func coinChange2(coins []int, amount int) int {
	n := len(coins)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0

	for i, num := range coins {
		for c := 0; c <= amount; c++ {
			if c < coins[i] {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-num]+1)
			}
		}
	}

	ans := f[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// 空间优化
func coinChange3(coins []int, amount int) int {
	n := len(coins)

	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0

	for i, num := range coins {
		for c := 0; c <= amount; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-num]+1)
			}
		}
	}
	ans := f[n%2][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, num := range coins {
		for c := num; c <= amount; c++ {
			f[c] = min(f[c], f[c-num]+1)
		}
	}
	ans := f[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
