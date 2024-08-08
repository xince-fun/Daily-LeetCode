package main

/*
capacity: 背包容量
w[i]: 第 i 个物品的体积
v[i]: 第 i 个物品的价值
返回： 所选物品体积和不超过 capacity 前提下，所能得到的最大价值和
*/

func zeroOneKnapSack(capacity int, w, v []int) int {
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
		memo[i][c] = max(dfs(i-1, c), dfs(i-1, c-w[i])+v[i])
		return memo[i][c]
	}
	return dfs(n-1, capacity)
}

func findTargetSumWays1(nums []int, target int) int {
	// 添加正号的 p
	// 负号的 sum - p
	// p - (sum - p) = t
	// 2p = sum + t
	// p = (sum + t) / 2
	// 变成"从nums中选择一些数字，使它们的和恰好等于(s+t)/2"的方案数
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

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
				return 1
			}
			return 0
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		if c < nums[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
		}
		return memo[i][c]
	}
	return dfs(n-1, target)
}

func findTargetSumWays2(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1

	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i][c-num]
			}
		}
	}
	return f[n][target]
}

// 空间优化
// 每次把f[i+1]算完之后，后面就不会用到f[i]了 每时每刻都只有两个数组中的元素在参与状态转移

// f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-w[i]]

func findTargetSumWays3(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2
	n := len(nums)
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1

	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-num]
			}
		}
	}
	return f[n%2][target]
}

func findTargetSumWays(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	f[0] = 1

	for _, num := range nums {
		for c := target; c >= num; c-- {
			f[c] = f[c] + f[c-num]
		}
	}
	return f[target]
}

// 如果是至多target
// 能递归到终点说明 c >= 0
func findTargetSumWays5(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	for i := range f {
		f[i] = 1
	}

	for _, num := range nums {
		for c := target; c >= num; c-- {
			f[c] = f[c] + f[c-num]
		}
	}
	return f[target]

	/*
		var dfs func(int, int) int
		dfs = func(i, c int) int {
			if i < 0 {
				return 1
			}
			if memo[i][c] != -1 {
				return memo[i][c]
			}
			if c < nums[i] {
				memo[i][c] = dfs(i-1, c)
			} else {
				memo[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
			}
			return memo[i][c]
		}
		return dfs(n-1, target)
	*/
}

// 如果是至少target
func findTargetSumWays6(nums []int, target int) int {
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	for i := range f {
		f[i] = 1
	}
	// f[0]表示至少为0的方案数，也就表示没有任何约束
	// 比如选第i个物品后 c <= 0了，就表示前面的 i-1 个物品可以不受约束的随意选或不选了
	for _, num := range nums {
		for c := target; c >= 0; c-- {
			f[c] = f[c] + f[max(c-num, 0)]
		}
	}
	return f[target]

	/*
		var dfs func(int, int) int
		dfs = func(i, c int) int {
			if i < 0 {
				if c <= 0 {
					return 1
				}
				return 0
			}
			if memo[i][c] != -1 {
				return memo[i][c]
			}
			memo[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])

			return memo[i][c]
		}
		return dfs(n-1, target)
	*/
}
