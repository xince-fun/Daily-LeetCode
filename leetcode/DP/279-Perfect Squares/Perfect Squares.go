package main

import "math"

// 先考虑记忆化搜索

func numSquares1(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	m := len(nums)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
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
		if c < nums[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = min(dfs(i-1, c), dfs(i-1, c-nums[i])+1)
		}
		return memo[i][c]
	}
	return dfs(m-1, n)
}

// 改dp
func numSquares2(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	m := len(nums)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	f[0][0] = 0
	for i, num := range nums {
		for c := 0; c <= n; c++ {
			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-num]+1)
			}
		}
	}
	return f[m][n]
}

func numSquares3(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	m := len(nums)

	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	f[0][0] = 0
	for i, num := range nums {
		for c := 0; c <= n; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-num]+1)
			}
		}
	}
	return f[m%2][n]
}

func numSquares(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
	}

	for _, num := range nums {
		for c := 0; c <= n; c++ {
			if c >= num {
				f[c] = min(f[c], f[c-num]+1)
			}
		}
	}
	return f[n]
}
