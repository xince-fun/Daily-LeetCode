package main

import "sort"

/*
不选nums[i]: f[i] = f[i-1]
选nums[i]: nums[i] <= mx: f[i] = f[i-2]+1

f[i] = max(f[i-1], f[i-2]+1)
*/

func minCapability(nums []int, k int) int {
	// 二分窃取能力
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}
