package main

import "fmt"

func rob1(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	// 0 不偷 1 偷
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func rob3(nums []int) int {
	cur, pre := 0, 0
	for _, num := range nums {
		cur, pre = max(pre+num, cur), cur
	}
	return cur
}

// 状态定义？状态转移方程？
// 启发思路： 选或不选/选哪个

// DP萌新三步
// 					  入参和返回值
// 思考回溯要怎么写  ->  递归到哪里
// 改成记忆化搜索        递归边界和入口
// 1:1 翻译成递推

func rob4(nums []int) int {
	n := len(nums)
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		return res
	}
	return dfs(n - 1)
}

func rob5(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		memo[i] = res
		return res
	}
	return dfs(n - 1)
}

// 自顶向下算 = 记忆化搜索
// 自底向上算 = 递推
//				dfs -> f数组
// 1:1翻译成递推 递归 -> 循环
//			   递归边界 -> 数组初始值

// dfs(i) = max(dfs(i-1), dfs(i-2)+nums[i])
// f[i] = max(f[i-1], f[i-2]+nums[i])
// f[i+2] = max(f[i+1], f[i]+nums[i])

func rob6(nums []int) int {
	n := len(nums)
	f := make([]int, n+2)
	for i, x := range nums {
		f[i+2] = max(f[i+1], f[i]+x)
	}
	return f[n+1]
}

// 当前=max(上一个，上上一个+nums[i])
// f0表示上上一个，f1表示上一个
// newF = max(f1,f0+nums[i])
// f0=f1
// f1=newF

func rob(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		newF := max(f1, f0+x)
		f0 = f1
		f1 = newF
	}
	// 最后一次算的newF
	return f1
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob1(nums))
}
