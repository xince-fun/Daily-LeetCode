package main

// 滑动窗口求子数组
// 符合的时候 ans += (right - left + 1)
// 不符合的时候 ans += left

func numSubarrayProductLessThanK1(nums []int, k int) (ans int) {
	total := 1
	left := 0
	for right, x := range nums {
		total *= x
		for total >= k && left < right {
			total /= nums[left]
			left++
		}
		if total < k {
			ans += (right - left + 1)
		}
	}
	return
}

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	total, left := 1, 0
	for right, x := range nums {
		total *= x
		for total >= k {
			total /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return
}
