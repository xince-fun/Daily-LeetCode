package main

// L-1始终是红色 R+1始终是蓝色
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target { // [mid+1, right]
			left = mid + 1
		} else { // [left, mid-1]
			right = mid - 1
		}
	}
	return left
}

func lowerBound2(nums []int, target int) int {
	left, right := 0, len(nums) // [left, right)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target { // [mid+1, right)
			left = mid + 1
		} else { // [left, mid)
			right = mid
		}
	}
	return left
}

func lowerBound3(nums []int, target int) int {
	left, right := -1, len(nums)-1 // (left, right)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target { // (mid, right)
			left = mid
		} else { // (left, mid)
			right = mid
		}
	}
	return left
}

// >= 既是这种
// > 转为  >= x+1
// < 转为 (>=x)-1
// <= 转为 (>x)-1

// 闭区间 染色法
func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}
