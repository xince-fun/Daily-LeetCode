package main

import "sort"

const mod = 1e9 + 7

func maxSumRangeQuery(nums []int, requests [][]int) (ans int) {
	n := len(nums)
	d := make([]int, n+1)
	for _, r := range requests {
		d[r[0]]++
		d[r[1]+1]--
	}
	for i := 1; i <= n; i++ {
		d[i] += d[i-1]
	}
	sort.Ints(d[:n])
	sort.Ints(nums)

	for i := n - 1; i >= 0; i-- {
		ans = (ans + nums[i]*d[i]) % mod
	}
	return ans
}
