package main

func lowerBound(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := right + (left-right)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func maximumCount(nums []int) int {
	// >0的idx
	start := lowerBound(nums, 0)
	// <0的idx
	end := lowerBound(nums, 1) - 1
	return max(end+1, len(nums)-start)
}
