package main

import "fmt"

func longestSubarray(nums []int) (ans int) {
	left := 0
	cnt := 0
	for right, x := range nums {
		if x == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left)
	}
	return
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
