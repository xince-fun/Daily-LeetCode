package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	st := make([]int, 0)
	mapp := make(map[int]int)
	for i := m - 1; i >= 0; i-- {
		x := nums2[i]
		for len(st) != 0 && st[len(st)-1] <= x {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			mapp[x] = -1
		} else {
			mapp[x] = st[len(st)-1]
		}
		st = append(st, x)
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = mapp[nums1[i]]
	}
	return ans
}
