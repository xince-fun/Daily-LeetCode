package main

import "math"

// 选一些数 sums = target 价值为 1
// capacity = target
// 返回最大的价值和

// 先用dfs()来完成

// dfs(i,c) = max(dfs(i-1,c),dfs(i-1,c-w[i])+1)
func lengthOfLongestSubsequence1(nums []int, target int) int {
	// 从nums中选一些数，使它们的和恰好等于 target的方案数
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MinInt
		}
		if memo[i][c] > -1 {
			return memo[i][c]
		}
		if c < nums[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = max(dfs(i-1, c), dfs(i-1, c-nums[i])+1)
		}
		return memo[i][c]
	}
	ans := dfs(n-1, target)
	if ans < 0 {
		return -1
	}
	return ans
}

func lengthOfLongestSubsequence2(nums []int, target int) int {
	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[0][0] = 0
	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = max(f[i][c], f[i][c-num]+1)
			}
		}
	}
	if f[n][target] > 0 {
		return f[n][target]
	}
	return -1
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	f := make([]int, target+1)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0
	for _, num := range nums {
		for c := target; c >= num; c-- {
			f[c] = max(f[c], f[c-num]+1)
		}
	}
	if f[target] > 0 {
		return f[target]
	}
	return -1
}
