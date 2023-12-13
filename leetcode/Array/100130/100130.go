package main

import "fmt"

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, n := range nums1 {
		m1[n]++
	}
	for _, n := range nums2 {
		m2[n]++
	}
	cnt1, cnt2 := 0, 0
	for k, v := range m1 {
		if m2[k] > 0 {
			cnt1 += v
		}
	}
	for k, v := range m2 {
		if m1[k] > 0 {
			cnt2 += v
		}
	}
	return []int{cnt1, cnt2}
}

func main() {
	nums1 := []int{4, 3, 2, 3, 1}
	nums2 := []int{2, 2, 5, 2, 3, 6}
	fmt.Println(findIntersectionValues(nums1, nums2))
}
