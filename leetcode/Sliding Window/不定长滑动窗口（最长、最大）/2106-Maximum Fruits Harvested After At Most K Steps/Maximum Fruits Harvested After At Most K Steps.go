package main

import "sort"

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	// 向左最远 能到的
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	right, s := left, 0
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1] // 从 fruits[left][0] 到 startPos的水果数
	}
	ans := s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1]
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1]
			left++
		}
		ans = max(ans, s)
	}
	return ans
}
