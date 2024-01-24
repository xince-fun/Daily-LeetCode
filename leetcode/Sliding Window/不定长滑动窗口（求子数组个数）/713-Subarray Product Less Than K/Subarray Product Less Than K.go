package main

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	product := 1
	left := 0
	for right, x := range nums {
		product *= x
		for product >= k {
			product /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return
}
