package main

import "math"

func maxSubArray1(nums []int) (ans int) {
	ans = math.MinInt
	count := 0
	for _, n := range nums {
		count += n
		if count > ans {
			ans = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return
}

func maxSubArray2(nums []int) (ans int) {
	ans = nums[0]
	count := 0
	for _, n := range nums {
		if count <= 0 {
			count = n
		} else {
			count += n
		}
		ans = max(ans, count)
	}
	return
}

func maxSubArray(nums []int) (ans int) {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	ans = dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		ans = max(ans, dp[i])
	}
	return
}

func maxSubArray3(nums []int) (ans int) {
	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxSubArraySum(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	return max(maxSubArraySum(nums, left, mid), maxSubArraySum(nums, mid+1, right), maxCrossingSum(nums, left, mid, right))
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	// 一定包含 nums[mid]
	leftSum, rightSum := math.MinInt, math.MinInt
	sum := 0
	// 左半边包含 nums[mid]
	// 计算以mid结尾的最大子数和
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}
