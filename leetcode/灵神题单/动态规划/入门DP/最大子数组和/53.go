package main

import (
	"math"
	"slices"
)

// 定义 f[i] 表示以 a[i] 结尾的最大子数组和， 不和左边拼起来就是 f[i] = a[i] , 和 i左边拼起来就是 f[i] = f[i-1] + a[i]
// 取最大值就得到了状态转移方程， f[i] = max(f[i-1],0) + a[i] 答案为 max(f)

func maxSubArray1(nums []int) (ans int) {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1], 0) + nums[i]
	}
	return slices.Max(f)
}

func maxSubArray(nums []int) (ans int) {
	ans = math.MinInt
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return
}

func maxSubArray2(nums []int) (ans int) {
	ans = math.MinInt
	minPreSum := 0
	preSum := 0
	for _, x := range nums {
		preSum += x
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return
}
