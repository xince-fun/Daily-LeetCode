package main

import (
	"fmt"
)

func maximumSetSize(nums1 []int, nums2 []int) int {
	set1 := map[int]bool{}
	for _, x := range nums1 {
		set1[x] = true
	}
	set2 := map[int]bool{}
	for _, x := range nums2 {
		set2[x] = true
	}
	common := 0
	for x := range set1 {
		if set2[x] {
			common++
		}
	}

	n1 := len(set1)
	n2 := len(set2)
	ans := n1 + n2 - common

	m := len(nums1) / 2
	if n1 > m {
		mn := min(n1-m, common)
		ans -= n1 - mn - m
		common -= mn
	}

	if n2 > m {
		n2 -= min(n2-m, common)
		ans -= n2 - m
	}

	return ans

}

func main() {
	nums1 := []int{1, 2, 1, 2}
	nums2 := []int{1, 1, 1, 1}
	fmt.Println(maximumSetSize(nums1, nums2))
}
