package main

import (
	"slices"
	"sort"
)

func maximumBeauty1(nums []int, k int) (ans int) {
	slices.Sort(nums)

	i, j := 0, 0
	for j < len(nums) {
		if nums[j] > nums[i]+2*k {
			i = sort.Search(1e5+1, func(b int) bool {
				return nums[b]+2*k >= nums[j]
			})
		}
		ans = max(ans, j-i+1)
		j++
	}
	return
}

func maximumBeauty2(nums []int, k int) (ans int) {
	slices.Sort(nums)

	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maximumBeauty(nums []int, k int) (ans int) {
	m := 0
	for _, num := range nums {
		m = max(m, num)
	}

	diff := make([]int, m+2)
	for _, x := range nums {
		diff[max(x-k, 0)]++
		diff[min(x+k+1, m+1)]--
	}
	count := 0
	for _, x := range diff {
		count += x
		ans = max(ans, count)
	}
	return
}
