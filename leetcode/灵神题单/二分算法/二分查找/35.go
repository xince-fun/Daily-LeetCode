package main

// >=
func lowerBound(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func searchInsert(nums []int, target int) int {
	return lowerBound(nums, target)
}
