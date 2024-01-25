package main

import (
	"fmt"
)

/*
对于 nums = [1, 2, 5, 9] thresold = 6

开始选择了 5，得到结果 5 < 6
选择4，结果为 1 + 1 + 2 + 3 = 7 > 6
选择 6，得到结果 1 + 1 + 1 + 2 = 5 < 6

所以对于 left-1 的回答一定是「否」
对于 right+1 的回答一定是「是」
*/

func smallestDivisor1(nums []int, threshold int) int {
	// 闭区间 [left, right]
	// 除数最大为nums[n-1]
	left := 1
	right := -1
	for _, num := range nums {
		right = max(right, num)
	}
	cal := func(div int) (ans int) {
		for _, num := range nums {
			ans += (num + div - 1) / div
		}
		return ans
	}
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「否」
		// right+1 的回答一定是「是」
		mid := (left + right) / 2
		if cal(mid) > threshold {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 循环结束后， right 等于 left-1，回答一定为「否」
	// 根据循环不变量， right+1 现在是最大的回答为「是」的数
	return right + 1
}

// 左闭右开
func smallestDivisor2(nums []int, threshold int) int {
	// 闭区间 [left, right)
	// 除数最大为nums[n-1]
	left := 1
	right := -1
	for _, num := range nums {
		right = max(right, num)
	}
	right = right + 1
	cal := func(div int) (ans int) {
		for _, num := range nums {
			ans += (num + div - 1) / div
		}
		return ans
	}
	for left < right {
		// 循环不变量：
		// left-1 的回答一定是「否」
		// right 的回答一定是「是」
		mid := (left + right) / 2
		if cal(mid) > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// 循环结束后， right 等于 left，回答一定为「否」
	// 根据循环不变量， right 现在是最大的回答为「是」的数
	return right
}

// 左开右闭
func smallestDivisor3(nums []int, threshold int) int {
	// 闭区间 (left, right]
	// 除数最大为nums[n-1]
	left := 0
	right := -1
	for _, num := range nums {
		right = max(right, num)
	}
	cal := func(div int) (ans int) {
		for _, num := range nums {
			ans += (num + div - 1) / div
		}
		return ans
	}
	for left < right {
		// 循环不变量：
		// left 的回答一定是「否」
		// right+1 的回答一定是「是」
		mid := (left + right + 1) / 2
		if cal(mid) > threshold {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 结束是 left == right
	// 根据循环不变量，right+1 现在是最大的回答为「是」的数
	return right + 1
}

// 开区间
func smallestDivisor(nums []int, threshold int) int {
	// 闭区间 (left, right)
	// 除数最大为nums[n-1]
	left := 0
	right := -1
	for _, num := range nums {
		right = max(right, num)
	}
	right = right + 1
	cal := func(div int) (ans int) {
		for _, num := range nums {
			ans += (num + div - 1) / div
		}
		return ans
	}
	for left+1 < right {
		// 循环不变量：
		// left 的回答一定是「否」
		// right 的回答一定是「是」
		mid := (left + right + 1) / 2
		if cal(mid) > threshold {
			left = mid
		} else {
			right = mid
		}
	}
	// 根据循环不变量，left+1 现在是最大的回答为「是」的数
	return right
}

func main() {
	nums := []int{44, 22, 33, 11, 1}
	threshold := 5
	fmt.Println(smallestDivisor(nums, threshold))
}
