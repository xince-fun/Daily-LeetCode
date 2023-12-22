package main

// 恰好capacity，求是否能

func canPartition1(nums []int) bool {
	// sum / 2
	// 转换为0-1背包问题
	// 变成从nums中选择一些数字，使它们的和恰好为sum/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	n := len(nums)
	memo := make([][]*bool, n)
	for i := range memo {
		memo[i] = make([]*bool, sum+1)
	}
	var dfs func(i, c int) bool
	dfs = func(i, c int) bool {
		if i < 0 {
			return c == 0
		}
		if memo[i][c] != nil {
			return *memo[i][c]
		}
		take := false
		if c >= nums[i] {
			take = dfs(i-1, c-nums[i])
		}
		dontTake := dfs(i-1, c)
		result := take || dontTake
		memo[i][c] = &result
		return result
	}
	return dfs(n-1, sum)
}

// dp

func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	n := len(nums)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, sum+1)
	}
	f[0][0] = true

	for i, num := range nums {
		for c := 0; c <= sum; c++ {
			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] || f[i][c-num]
			}
		}
	}
	return f[n][sum]
}

func canPartition3(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	n := len(nums)
	f := make([][]bool, 2)
	for i := range f {
		f[i] = make([]bool, sum+1)
	}
	f[0][0] = true

	for i, num := range nums {
		for c := 0; c <= sum; c++ {
			if c < num {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] || f[i%2][c-num]
			}
		}
	}
	return f[n%2][sum]
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	f := make([]bool, sum+1)
	f[0] = true

	for _, num := range nums {
		for c := sum; c >= 0; c-- {
			if c >= num {
				f[c] = f[c] || f[c-num]
			}
		}
	}
	return f[sum]
}
