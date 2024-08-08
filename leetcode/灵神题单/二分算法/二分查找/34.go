package main

import "sort"

// >= 既是这种
// > 转为  >= x+1
// < 转为 (>=x)-1
// <= 转为 (>x)-1

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

func maximumCount1(nums []int) int {
	// >0的idx
	start := lowerBound(nums, 1)
	// >-1的idx
	end := lowerBound(nums, 0)
	return max(end, len(nums)-start)
}

// sort.SearchInts 返回的是 >= x 的第一个位置
// >= x 就是这种情况
// > 转换为 >= (x+1)
// < 转换为 (>=x)-1
// <= 转换为 (>x)-1 即 >=(x+1)-1

func maximumCount(nums []int) int {
	// >-1的第一个位置
	neg := sort.SearchInts(nums, 0)
	pos := len(nums) - sort.SearchInts(nums, 1)
	return max(neg, pos)
}
