package main

// 小于最后一个数，就是蓝色，否则就是红色
func findMin(nums []int) int {
	// (left, right)
	left, right := -1, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	return nums[right]
}
