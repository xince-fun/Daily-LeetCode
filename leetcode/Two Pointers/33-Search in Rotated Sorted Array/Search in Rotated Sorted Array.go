package main

func search(nums []int, target int) int {
	// (left, right)
	var is_blue func(int) bool
	is_blue = func(i int) bool {
		end := nums[len(nums)-1]
		if nums[i] > end {
			return target > end && nums[i] >= target
		} else {
			return target > end || nums[i] >= target
		}
	}
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if is_blue(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}
