package main

func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := -1, n-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
