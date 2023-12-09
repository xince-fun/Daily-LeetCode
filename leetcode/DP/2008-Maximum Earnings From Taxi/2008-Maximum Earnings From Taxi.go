package main

import (
	"fmt"
)

// dp[1] = 0 答案为dp[n]
// 如果不在i下车, dp[i] = dp[i-1]
// 如果在i下车, dp[i] = max(dp[i-1],dp[starti]+i-starti+tipi)

// 原始递归 超时
func maxTaxiEarnings1(n int, rides [][]int) int64 {
	type pair struct {
		s, t int
	}
	groups := make([][]pair, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], pair{start, end - start + tip})
	}

	var dfs func(int) int64
	dfs = func(i int) int64 {
		if i == 0 {
			return 0
		}
		res := dfs(i - 1)
		for _, p := range groups[i] {
			res = max(res, dfs(p.s)+int64(p.t))
		}
		return res
	}
	return dfs(n)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// 记忆化
func maxTaxiEarnings2(n int, rides [][]int) int64 {
	type pair struct {
		s, t int
	}
	groups := make([][]pair, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], pair{start, end - start + tip})
	}
	memo := make([]int64, n+1)
	for i := range memo {
		memo[i] = -1 // 没搜索过
	}
	var dfs func(int) int64
	dfs = func(i int) int64 {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		res := dfs(i - 1)
		for _, p := range groups[i] {
			res = max(res, dfs(p.s)+int64(p.t))
		}
		*p = res
		return res
	}
	return dfs(n)
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	type pair struct {
		s, t int
	}
	groups := make([][]pair, n+1)
	for _, r := range rides {
		start, end, tip := r[0], r[1], r[2]
		groups[end] = append(groups[end], pair{start, end - start + tip})
	}
	f := make([]int64, n+1)
	for i := 2; i <= n; i++ {
		f[i] = f[i-1]
		for _, p := range groups[i] {
			f[i] = max(f[i], f[p.s]+int64(p.t))
		}
	}
	return f[n]
}

func main() {

	n := 20
	rides := [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}

	fmt.Println(maxTaxiEarnings(n, rides))
}
