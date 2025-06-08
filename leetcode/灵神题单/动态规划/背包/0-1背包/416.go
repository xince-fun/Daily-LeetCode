package main

// nums = [1,5,11,5]
// 选出一些数， 使总和为 sum/2

func canPartition1(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2

	n := len(nums)
	memo := make([][]*bool, n)
	for i := range memo {
		memo[i] = make([]*bool, target+1)
	}
	var dfs func(int, int) bool
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
	return dfs(n-1, target)
}

func canPartition2(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2

	n := len(nums)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, target+1)
	}
	f[0][0] = true

	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] || f[i][c-num]
			}
		}
	}
	return f[n][target]
}

func canPartition(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 != 0 {
		return false
	}
	target /= 2

	f := make([]bool, target+1)
	f[0] = true

	for _, num := range nums {
		for c := target; c >= num; c-- {
			f[c] = f[c] || f[c-num]
		}
	}
	return f[target]
}
