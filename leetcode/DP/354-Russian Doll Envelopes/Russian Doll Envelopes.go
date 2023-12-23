package main

import (
	"sort"
)

func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	n := len(envelopes)

	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		res := 0
		if memo[i] > 0 {
			return memo[i]
		}
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				res = max(res, dfs(j))
			}
		}
		memo[i] = res + 1
		return memo[i]
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

func maxEnvelopes2(envelopes [][]int) (ans int) {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	n := len(envelopes)

	f := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				f[i] = max(f[i], f[j])
			}
		}
		f[i] += 1
	}
	for i := 0; i < n; i++ {
		if f[i] > ans {
			ans = f[i]
		}
	}
	return
}

func maxEnvelopes(envelopes [][]int) (ans int) {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	g := []int{}
	for _, num := range envelopes {
		j := sort.SearchInts(g, num[1])
		if j == len(g) {
			g = append(g, num[1])
		} else {
			g[j] = num[1]
		}
	}
	return len(g)
}
