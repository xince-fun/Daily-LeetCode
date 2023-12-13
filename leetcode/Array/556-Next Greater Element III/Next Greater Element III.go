package main

import (
	"math"
	"sort"
)

func nextGreaterElement1(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	length := len(digits)
	index := -1
	for i := 0; i < length-1 && index == -1; i++ {
		if digits[i+1] < digits[i] {
			index = i + 1
		}
	}
	if index == -1 {
		return -1
	}
	for i := 0; i < index; i++ {
		if digits[i] > digits[index] {
			digits[i], digits[index] = digits[index], digits[i]
			break
		}
	}
	for l, r := 0, index-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	var ans int64
	for i := length - 1; i >= 0; i-- {
		ans = ans*10 + int64(digits[i])
	}
	if ans > math.MaxInt32 {
		return -1
	}
	return int(ans)
}

func nextGreaterElement(n int) int {
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	stack := []int{}

	x, y := -1, -1
	for i, num := range nums {
		for len(stack) > 0 && nums[len(stack)-1] > num {
			x = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if x >= 0 {
			y = i
			break
		}
		stack = append(stack, i)
	}
	if x == -1 {
		return -1
	}
	nums[x], nums[y] = nums[y], nums[x]
	var res int
	for i := len(nums) - 1; i >= y; i-- {
		res = res*10 + nums[i]
	}
	nums = nums[:y]
	sort.Ints(nums)
	for _, num := range nums {
		res = res*10 + num
	}
	if res > math.MaxInt32 {
		return -1
	}
	return res
}
