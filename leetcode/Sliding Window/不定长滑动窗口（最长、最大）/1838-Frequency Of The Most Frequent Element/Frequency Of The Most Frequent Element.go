package main

import (
	"cmp"
	"fmt"
	"slices"
)

func maxFrequency(nums []int, k int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	left := 0
	sum := 0
	for right, v := range nums {
		sum += v
		for sum+k < v*(right-left+1) {
			sum -= nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func main() {
	nums := []int{1, 4, 8, 13}
	k := 5
	fmt.Println(maxFrequency(nums, k))
}
