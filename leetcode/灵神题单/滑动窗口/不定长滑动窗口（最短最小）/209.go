package main

func minSubArrayLen(target int, nums []int) (ans int) {
	ans = len(nums) + 1
	sum := 0
	left := 0
	for right, num := range nums {
		sum += num
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans <= len(nums) {
		return ans
	}
	return 0
}
