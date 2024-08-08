package main

import (
	"slices"
	"sort"
)

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	mx := slices.Min(stock)
	mx += budget
	for _, comp := range composition {
		// 二分 <= budget 的最大份数
		// 先找 > budget 的第一个，再 -1
		cnt := sort.Search(mx, func(mid int) bool {
			sum := 0
			mid++
			for i, c := range comp {
				sum += max(0, mid*c-stock[i]) * cost[i]
			}
			return sum > budget
		})
		ans = max(ans, cnt)
	}
	return
}
