package main

// 当前操作？ 爬 1或 2 阶楼梯
// 子问题？ 到达 第 i 阶楼梯时的方法数
// 下一个问题？ 分类讨论：
// 爬1阶：从前i-1阶得到的方法数
// 爬2阶：从前i-2阶得到的方法数

// res = dfs(i-1) + dfs(i-2)
// dfs(0) = dfs(1) = 1

func climbStairs1(n int) int {
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		res := dfs(i-1) + dfs(i-2)
		return res
	}
	return dfs(n)
}

func climbStairs2(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := dfs(i-1) + dfs(i-2)
		memo[i] = res
		return res
	}
	return dfs(n)
}

// 递归到递推

// dfs(i) = dfs(i-1) + dfs(i-2)
// f[i] = f[i-1] + f[i-2]
// f[i+2] = f[i+1] + f[i]

func climbStairs3(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// newF = f_0 + f_1
// 对于下一轮
// 「上上一个状态」 就是f_1, f_0 = f_1
// 「上一个状态」就是 newF, f1 = newF

// 空间优化
func climbStairs(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
