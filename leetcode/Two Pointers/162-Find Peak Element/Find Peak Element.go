package main

func findPeakElement(nums []int) int {
	// (-1, n-1)
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	return right
}
