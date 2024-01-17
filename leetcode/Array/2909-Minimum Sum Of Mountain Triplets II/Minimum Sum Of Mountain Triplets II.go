package main

import (
	"fmt"
	"math"
)

func minimumSum1(nums []int) int {
	n := len(nums)
	// 前后缀分解
	// 对于 一个 i， 左侧最小的数， 右侧最小的数
	// pre[1]左侧的数，
	preMin := make([]int, n)
	sufMin := make([]int, n)
	// 第一个数左边没有最小的数，初始化为MaxInt
	preMin[0] = math.MaxInt
	for i := 1; i < n; i++ {
		preMin[i] = min(preMin[i-1], nums[i-1])
	}
	sufMin[n-1] = math.MaxInt
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i+1])
	}
	fmt.Println(preMin)
	fmt.Println(sufMin)
	ans := math.MaxInt
	for i := 1; i < n-1; i++ {
		if nums[i] > preMin[i] && nums[i] > sufMin[i] {
			ans = min(ans, nums[i]+preMin[i]+sufMin[i])
		}
	}
	if ans >= math.MaxInt {
		return -1
	}
	return ans
}

// 灵神的
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}
	ans := math.MaxInt
	pre := nums[0]
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] {
			ans = min(ans, pre+nums[j]+suf[j+1])
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func main() {
	nums := []int{6, 5, 4, 3, 4, 5}
	fmt.Println(minimumSum(nums))
}
