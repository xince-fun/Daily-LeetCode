package main

func longestSubarray(nums []int) (ans int) {
	count := 0
	left := 0
	for right, num := range nums {
		count += num ^ 1
		for count > 1 {
			count -= nums[left] ^ 1
			left++
		}
		ans = max(ans, right-left)
	}
	return
}
