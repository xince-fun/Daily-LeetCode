package main

func numSubarrayProductLessThanK1(nums []int, k int) int {
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

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans := 0

	left, product := 0, 1

	for right, x := range nums {
		product *= x
		for product >= k {
			product /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
