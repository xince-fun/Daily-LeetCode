package main

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	ans, n := len(nums)+1, len(nums)
	left, right := 0, 0
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
