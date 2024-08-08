package main

// 选不选的问题

// 选择下一个  +nums[j] -x
// 不选下一个  +nums[j]
// 什么都不选

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	dp := make([][2]int, n+1)

	for i := n - 1; i >= 0; i-- {
		x1 := nums[i] % 2
		dp[i][x1^1] = dp[i+1][x1^1]
		dp[i][x1] = max(dp[i+1][x1], dp[i+1][x1^1]-x) + nums[i]
	}

	return int64(dp[0][nums[0]%2])
}
