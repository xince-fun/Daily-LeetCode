package main

// 当前操作？ 枚举 **第** i 个房子选/不选
// 子问题？ 从 **前** i 个房子中得到的最大金额和
// 下一个问题？ 分类讨论：
// 不选： 从 **前** i-1 个房子中得到的最大金额和
//  选： 从 **前** i-2 个房子中得到的最大金额和

func rob1(nums []int) int {
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

func rob2(nums []int) int {
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

func rob3(nums []int) int {
	n := len(nums)
	f := make([]int, n+2)
	for i, x := range nums {
		f[i+2] = max(f[i+1], f[i]+x)
	}
	return f[n+1]
}

func rob4(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		newF := max(f1, f0+x)
		f0 = f1
		f1 = newF
	}
	return f1
}

func rob(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		f0, f1 = f1, max(f1, f0+x)
	}
	return f1
}
