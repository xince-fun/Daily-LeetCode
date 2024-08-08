package main

// 当前操作？爬 1或 2 阶楼梯
// 子问题？ 到达第 i 阶楼梯时的最小费用
// 下一个问题？分类讨论：
// 爬1阶楼梯？从前i-1阶楼梯得到的最小费用
// 爬2阶楼梯？从前i-2阶楼梯得到的最小费用

// dfs(i) = min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
// dfs(0) = 0, dfs(1) = 1, 一开始就在 0 或 1，不需要花费

func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		return res
	}
	return dfs(n)
}

func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}
		if memo[i+1] != -1 {
			return memo[i]
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		memo[i] = res
		return res
	}
	return dfs(n)
}

// 翻译为递推
// f[i] = min(f[i-1]+cost[i-1],f[i-2]+cost[i-2])

func minCostClimbingStairs3(cost []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 2; i <= n; i++ {
		f[i] = min(f[i-1]+cost[i-1], f[i-2]+cost[i-2])
	}
	return f[n]
}

func minCostClimbingStairs(cost []int) int {
	f0, f1 := 0, 0
	for i := 1; i < len(cost); i++ {
		f0, f1 = f1, min(f1+cost[i], f0+cost[i-1])
	}
	return f1
}
