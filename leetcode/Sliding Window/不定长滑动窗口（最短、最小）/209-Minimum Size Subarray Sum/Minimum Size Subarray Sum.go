package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) (ans int) {
	ans, n := len(nums)+1, len(nums)
	sum := 0
	left := 0
	for right, x := range nums {
		sum += x
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	fmt.Println(minSubArrayLen(target, nums))
}
