package main

import "math/bits"

// bitmask
func countArrangement1(n int) int {
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(s int) int {
		if s == 1<<n-1 {
			return 1
		}
		p := &memo[s]
		if *p != -1 {
			return *p
		}
		res := 0
		i := bits.OnesCount(uint(s)) + 1
		for j := 1; j <= n; j++ {
			if s>>(j-1)&1 == 0 && (i%j == 0 || j%i == 0) {
				res += dfs(s | 1<<(j-1))
			}
		}
		*p = res
		return res
	}
	return dfs(0)
}

func countArrangement(n int) int {
	f := make([]int, 1<<n)
	f[0] = 1
	for s := 1; s < 1<<n; s++ {
		i := bits.OnesCount(uint(s))
		for j := 1; j <= n; j++ {
			if s>>(j-1)&1 != 0 && (i%j == 0 || j%i == 0) {
				f[s] += f[s^1<<(j-1)]
			}
		}
	}
	return f[1<<n-1]
}
