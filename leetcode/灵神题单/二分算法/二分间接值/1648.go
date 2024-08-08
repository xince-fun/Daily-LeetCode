package main

import "sort"

// 2 8 4 10 6

// 8 7 6 5 4 3

func maxProfit(inventory []int, orders int) (ans int) {
	// 找到最大的下界 i，使减少的数量 >= orders
	// 随着 i 变大， 由 true 为 false
	// 所以先找 一个 i, < orders
	const mod = 1e9 + 7
	idx := sort.Search(1_000_000_001, func(k int) bool {
		cnt := 0
		for _, inv := range inventory {
			l := max(0, inv-k+1)
			cnt += l
		}
		return cnt <= orders
	})
	sum := 0
	cnt := 0

	for _, inv := range inventory {
		l := max(0, inv-idx+1)
		cnt += l
		sum += (inv + idx) * l / 2
		sum %= mod
	}
	if cnt < orders {
		sum += (orders - cnt) * (idx - 1)
		sum %= mod
	}
	return sum % (1e9 + 7)
}
