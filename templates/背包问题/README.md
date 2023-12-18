
## 0-1背包

```go
// 背包问题是顺序无关的

// 0-1背包：有n个物品，第i个物品的体积为w[i],价值为v[i]
// 每个物品至多选一个，求体积和不超过 capacity时的最大价值和

// 回溯三问

// 1.当前操作? 枚举第i个物品选或不选：
//	 不选，剩余容量不变；选，剩余容量减少w[i]
//
// 2.子问题？在剩余容量为c时，
//	 从前i个物品中得到的最大价值和
//
// 3.下一个子问题？分类讨论：
//   不选：在剩余容量为c时，
//        从前i-1个物品中得到的最大价值和
//   选：在剩余容量c-w[i]时，
//       从前i-1个物品中得到的最大价值和

// dfs(i,c) = max(dfs(i-1,c), dfs(i-1,c-w[i]) + v[i])

// 常见变形
// 1. 至多装capacity, 求方案数/最大价值和
// 2. 恰好装capacity, 求方案数/最大/最小价值和
// 3. 至少装capacity, 求方案数/最小价值和

// 加法原理，做一件事情有两种方式，方案数就是这两种的方案之和
// dfs(i,c) = dfs(i-1,c) + dfs(i-1,c-w[i])

// f[i][c] = f[i-1][c] + f[i-1][c-w[i]]
// f[i+1][c] = f[i][c] + f[i][c-w[i]]

func zeroOneKnapsack(capacity int, w, v []int) int {
	n := len(w)

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			return 0
		}
		if c < w[i] {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i-1, c-w[i])+v[i])
	}
	return dfs(n-1, capacity)
}

// 改为记忆化搜索
func zeroOneKnapsack1(capacity int, w, v []int) int {
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
		} else {
			memo[i][c] = max(dfs(i-1, c), dfs(i-1, c-w[i])+v[i])
		}
		return memo[i][c]
	}
	return dfs(n-1, capacity)
}

func findTargetSumWays1(nums []int, target int) int {
	// 假设添加正号的数和为p
	// 添加负数的数和为s-p
	// p-(s-p) = t
	// 2p = s+t
	// p = (s+t)/2
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
	// 假设添加正号的数和为p
	// 添加负数的数和为s-p
	// p-(s-p) = t
	// 2p = s+t
	// p = (s+t)/2
	// 变成"从nums中选择一些数字，使它们的和恰好等于(s+t)/2"的方案数
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2

	n := len(nums)
	f := make([][]int, n)
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
	// 假设添加正号的数和为p
	// 添加负数的数和为s-p
	// p-(s-p) = t
	// 2p = s+t
	// p = (s+t)/2
	// 变成"从nums中选择一些数字，使它们的和恰好等于(s+t)/2"的方案数
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
	return f[n][target]
}

// 进一步优化
// f[c] = f[c] + f[c-w[i]]

func findTargetSumWays(nums []int, target int) int {
	// 假设添加正号的数和为p
	// 添加负数的数和为s-p
	// p-(s-p) = t
	// 2p = s+t
	// p = (s+t)/2
	// 变成"从nums中选择一些数字，使它们的和恰好等于(s+t)/2"的方案数
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
			if c >= num {
				f[c] = f[c] + f[c-num]
			}
		}
	}
	return f[target]
}
```

## 完全背包

```go
// 完全背包： 有n种物品，第i种物品的体积为w[i],价值为v[i]
// 每个物品无限次重复选，求体积和不超过capcity时的最大价值和

// 回溯三问

// 1.当前操作？ 枚举第i个物品选或不选：
//		不选，剩余容量不变；选，剩余容量减少w[i]

// 2.子问题？ 在剩余容量为c时，
//		从前i种物品中得到的最大价值和

// 3. 下一个子问题？ 分类讨论：
//	  不选： 在剩余容量为c时，
//			从前i-1种物品中得到的最大价值和
//	  选：在剩余容量为c-w[i]时，
//		 从前i种物品中得到的最大价值和

// dfs(i,c) = max(dfs(i-1),c dfs(i,c-w[i]) + v[i])

// 常见变型：
// 1. 至多装capacity, 求方案数/最大价值和
// 2. 恰好装capacity, 求方案数/最大/最小价值和
// 3. 至少装capacity, 求方案数/最小价值和

// dfs(i,c) = min(dfs(i-1,c),dfs(i,c-w[i])+v[i])

// f[i][c] = min(f[i-1][c], f[i][c-w[i]]+v[i])

func unboundde_knapsack1(capacity int, w, v []int) int {
	n := len(w)

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			return 0
		}
		if c < w[i] {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i, c-w[i])+v[i])
	}
	return dfs(n-1, capacity)
}

// 记忆化搜索
func unboundde_knapsack2(capacity int, w, v []int) int {
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
		} else {
			memo[i][c] = max(dfs(i-1, c), dfs(i, c-w[i])+v[i])
		}
		return memo[i][c]
	}
	return dfs(n-1, capacity)
}

func coinChange1(coins []int, amount int) int {
	n := len(coins)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			memo[i][j] = -1 // 除 2 防止下面+1溢出
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
		if c < coins[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = min(dfs(i-1, c), dfs(i, c-coins[i])+1)
		}
		return memo[i][c]
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// 动态规划
func coinChange2(coins []int, amount int) int {
	n := len(coins)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0

	for i, num := range coins {
		for c := 0; c <= amount; c++ {
			if c < coins[i] {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-num]+1)
			}
		}
	}
	ans := f[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// 空间优化
func coinChange3(coins []int, amount int) int {
	n := len(coins)

	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0

	for i, num := range coins {
		for c := 0; c <= amount; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-num]+1)
			}
		}
	}
	ans := f[n%2][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// f[i][c] = min(f[i-1][c], f[i][c-w[i]]+v[i])
// f[i+1][c] = min(f[i][c], f[i+1][c-w[i]]+v[i])

// f[c] = min(f[c], f[c-w[i]]+v[i])

// 进一步优化
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, num := range coins {
		for c := 0; c <= amount; c++ {
			if c >= num {
				f[c] = min(f[c], f[c-num]+1)

			}
		}
	}
	ans := f[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

```

求方案数，
恰好的例子： 就是上面的例子
**如果是至多**： 
```go
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
```
改成
能递归到终点说明 c >= 0
```go
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
```

f 数组的初始化都改为 1

**至少的情况**

```go
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
		if c < nums[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
		}
		return memo[i][c]
	}
```

- c 循环从 c >= num 改成 c >= 0
- f[c-num] 改成 f[max(c-num, 0)]
表示把所有 c <=0 的状态都记录在 f[0] 中
```go
	for _, num := range nums {
		for c := target; c >= 0; c-- {
			if c >= num {
				f[c] = f[c] + f[c-num]
			}
		}
	}
```

![image-20231217233425322](https://raw.githubusercontent.com/xince-fun/Picgo/main/carl/image-20231217233425322.png)

![image-20231217233041767](https://raw.githubusercontent.com/xince-fun/Picgo/main/carl/image-20231217233041767.png)

![image-20231217233153976](https://raw.githubusercontent.com/xince-fun/Picgo/main/carl/image-20231217233153976.png)

