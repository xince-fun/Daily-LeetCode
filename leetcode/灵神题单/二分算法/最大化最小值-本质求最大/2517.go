package main

import (
	"slices"
	"sort"
)

// [13,5,1,8,21,2]
// [1,2,5,8,13,21]

//f(d_0) >= k且 f(d_0+1)<k

func maximumTastiness(price []int, k int) int {
	slices.Sort(price)
	// 随着 d 变大，越难满足
	// 二分 d, 求 < k 的第一个
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		cnt, pre := 1, price[0]
		for _, p := range price[1:] {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt < k
	})
}
