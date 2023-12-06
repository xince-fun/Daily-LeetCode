package main

import (
	"fmt"
	"sort"
)

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := make(map[float64]int)
	for len(nums) > 0 {
		small, big := nums[0], nums[len(nums)-1]
		nums = nums[1:]
		nums = nums[:len(nums)-1]
		avg := float64(small+big) / 2.0
		m[float64(avg)]++
	}
	return len(m)
}

func main() {
	nums := []int{9, 5, 7, 8, 7, 9, 8, 2, 0, 7}
	fmt.Println(distinctAverages(nums))
}
