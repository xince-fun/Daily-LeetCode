package main

import "slices"

// [1,2,3,4,5,6]
// rewardValues = [1,1,3,3]

// 初始为x, 如果 rewardValues[i] > x, x = x + rewardValues[i] 求 x 的最大值

// f[i][j] 表示能否从前 i 个数中， 选出和为 j 的子序列

// f[i][j] = f[i-1][j] || f[i-1][j-v]

func maxTotalReward(rewardValues []int) int {
	slices.Sort(rewardValues)
	rewardValues = slices.Compact(rewardValues)
	n := len(rewardValues)
	target := rewardValues[n-1]
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, target+1)
	}
	f[0][0] = true
	for i, num := range rewardValues {
		for c := 0; c <= target; c++ {
			f[i+1][c] = f[i][c] || f[i][c-num]
		}
	}
	for j := target; j >= 0; j-- {
		if f[n][j] {
			return j
		}
	}
	return 0
}
