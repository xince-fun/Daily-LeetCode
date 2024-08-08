package main

import "sort"

func maxValue1(n int, index int, maxSum int) int {
	// 随着 nums[index] 变大，先 true 后 false 可以二分nums[index]的值
	left, right := 0, maxSum+1

	getSum := func(mid, length int) int {
		if length == 0 {
			return 0
		}
		if mid > length+1 {
			return (2*mid - length - 1) * length / 2
		}
		return (mid-1)*mid/2 + length - mid + 1
	}

	check := func(mid int) bool {
		leftNum := index
		rightNum := n - index - 1
		return mid+getSum(mid, leftNum)+getSum(mid, rightNum) <= maxSum
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func maxValue(n int, index int, maxSum int) int {
	getSum := func(mid, length int) int {
		if length == 0 {
			return 0
		}
		if mid > length+1 {
			return (2*mid - length - 1) * length / 2
		}
		return (mid-1)*mid/2 + length - mid + 1
	}
	// 考虑二分 !f(x)
	return sort.Search(maxSum+1, func(mid int) bool {
		left := index
		right := n - index - 1
		sum := mid + getSum(mid, left) + getSum(mid, right)
		return sum > maxSum
	}) - 1
}
