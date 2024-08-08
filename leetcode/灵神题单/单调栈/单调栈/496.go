package main

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	st := []int{}
	for i, x := range nums2 {
		for len(st) > 0 && x > nums2[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			m[nums2[idx]] = x
		}
		st = append(st, i)
	}
	ans := make([]int, len(nums1))
	for i, x := range nums1 {
		if m[x] == 0 {
			ans[i] = -1
		} else {
			ans[i] = m[x]
		}
	}
	return ans
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	idx := make(map[int]int, len(nums1))
	for i, x := range nums1 {
		idx[x] = i
	}
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}

	st := []int{}

	for i, x := range nums2 {
		for len(st) > 0 && x > nums2[st[len(st)-1]] {
			ans[idx[nums2[st[len(st)-1]]]] = x
			st = st[:len(st)-1]
		}
		if _, ok := idx[x]; ok {
			st = append(st, i)
		}
	}
	return ans
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	idx := make(map[int]int, len(nums1))
	for i, x := range nums1 {
		idx[x] = i
	}
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		x := nums2[i]
		for len(st) > 0 && x >= nums2[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			if j, ok := idx[x]; ok {
				ans[j] = nums2[st[len(st)-1]]
			}
		}
		st = append(st, i)
	}

	return ans
}
