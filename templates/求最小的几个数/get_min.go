package main

import (
	"fmt"
)

// O(n) 的时间复杂度
func get2Min(nums []int) (int, int) {
	// 初始化两个比较大的数字 滚动
	a, b := 1001, 1001
	for _, x := range nums {
		if x < a {
			a, b = x, a
		} else if x < b {
			b = x
		}
	}
	return a, b
}

func get3Min(nums []int) (int, int, int) {
	// 初始三个
	a, b, c := 1001, 1001, 1001
	for _, x := range nums {
		if x < a {
			a, b = x, a
		} else if x < b {
			b, c = x, b
		} else if x < c {
			c = x
		}
	}
	return a, b, c
}

func main() {
	nums := []int{2, 4, 56, 57, 43, 1, 3}
	fmt.Println(get2Min(nums))
	fmt.Println(get3Min(nums))
}
