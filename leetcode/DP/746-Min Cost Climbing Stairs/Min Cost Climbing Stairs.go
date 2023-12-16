package main

import (
	"fmt"
)

// 递归回溯怎么写？
// dfs(i) = min(dfs(i-2)+cost[i-2],dfs(i-1)+cost[i-1])
// 如果最后爬了1个台阶，得先爬到i-1
// 如果最后爬了2个台阶，得先爬到i-2
// dfs(i) = min(dfs(i-1)+cost[i-1], dfs(i-2)))

// 会超时的
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)

	var dfs func(int) int
	dfs = func(i int) int {
		// 边界
		if i <= 1 {
			return 0
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		return res
	}
	return dfs(n)
}

// 记忆化搜索
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		memo[i] = res
		return res
	}
	return dfs(n)
}

// 动态规划
func minCostClimbingStairs3(cost []int) int {
	n := len(cost)
	f := make([]int, n+1)
	f[0], f[1] = 0, 0
	for i := 2; i <= n; i++ {
		f[i] = min(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[n]
}

// 空间优化
// 当前 = min(上上一个+cost[i-2],上一个+cost[i-1])
// f0表示上上一个，f1表示上一个
// newF = min(f0+cost[i-2], f1+cost[i-1])

func minCostClimbingStairs(cost []int) int {
	f0, f1 := 0, 0
	for i := 2; i <= len(cost); i++ {
		f0, f1 = f1, min(f0+cost[i-2], f1+cost[i-1])
	}
	return f1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
