package main

import (
	"fmt"
)

func countSubarrays(nums []int, k int) int64 {
	var ans int64
	n := len(nums)
	maxNum := 0
	for _, n := range nums {
		maxNum = max(maxNum, n)
	}
	start, end := 0, 0
	var count int

	for end < n {
		if nums[end] == maxNum {
			count++
		}
		if count >= k {
			for count >= k {
				ans += int64(n - end)
				if nums[start] == maxNum {
					count--
				}
				start++
			}
		}
		end++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 2, 3, 3}
	k := 2
	fmt.Println(countSubarrays(nums, k))
}
