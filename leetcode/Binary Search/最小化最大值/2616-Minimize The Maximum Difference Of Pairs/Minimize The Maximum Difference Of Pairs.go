package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	//
	return sort.Search(1e9, func(d int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= d {
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}

func main() {
	nums := []int{10, 1, 2, 7, 1, 3}
	p := 2
	fmt.Println(minimizeMax(nums, p))
}
