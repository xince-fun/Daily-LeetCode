package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, n := 0, len(nums)
	left, right := 0, 0
	product := 1

	for right < n {
		for product < k && right < n {
			product *= nums[right]
			right++
		}
		ans += 1<<(right-left) - 1
		for product >= k && left <= right {
			product /= nums[left]
			left++
		}
	}
	return ans
}
