package main

import "math"

// 当前操作？
// 子问题？ 在i的位置
// 下一个子问题？ 在i-1的位置怎么做操作

// dfs(i, j) 表示在第i+1天 完成 job[0]-job[j]的工作

func minDifficulty1(jobDifficulty []int, d int) (ans int) {
	n := len(jobDifficulty)

	memo := make([][]int, d)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		// 最后一天了必须全做
		if i == 0 {
			for _, x := range jobDifficulty[:j+1] {
				res = max(res, x)
			}
			return
		}
		res = math.MaxInt
		mx := 0
		for k := j; k >= i; k-- {
			mx = max(mx, jobDifficulty[k])
			res = min(res, dfs(i-1, k-1)+mx)
		}
		return
	}
	return dfs(d-1, n-1)
}

func minDifficulty(jobDifficulty []int, d int) (ans int) {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	f := make([][]int, d)
	f[0] = make([]int, n)
	f[0][0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		f[0][i] = max(f[0][i-1], jobDifficulty[i])
	}
	for i := 1; i < d; i++ {
		f[i] = make([]int, n)
		for j := n - 1; j >= i; j-- {
			f[i][j] = math.MaxInt
			mx := 0
			for k := j; k >= i; k-- {
				mx = max(mx, jobDifficulty[k])
				f[i][j] = min(f[i][j], f[i-1][k-1]+mx)
			}
		}
	}
	return f[d-1][n-1]
}
