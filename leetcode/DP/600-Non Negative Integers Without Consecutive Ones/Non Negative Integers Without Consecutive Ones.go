package main

import (
	"strconv"
)

func findIntegers(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	m := len(s)
	dp := make([][2]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(int, int, bool) int
	f = func(i, pre int, isLimit bool) (res int) {
		if i == m {
			return 1
		}
		if !isLimit {
			p := &dp[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		up := 1
		if isLimit {
			up = int(s[i] & 1)
		}
		res = f(i+1, 0, isLimit && up == 0)
		if pre == 0 && up == 1 {
			res += f(i+1, 1, isLimit)
		}
		return
	}
	return f(0, 0, true)
}
