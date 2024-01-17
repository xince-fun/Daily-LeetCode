package main

import "strings"

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	n := len(num2)
	const mod = 1_000_000_007
	num1 = strings.Repeat("0", n-len(num1)) + num1 // 补充前导零，与 num2对齐
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, min(9*n, max_sum)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, sum int, limitLow, limitHigh bool) (res int) {
		if sum > max_sum {
			return
		}
		if i == n {
			if sum >= min_sum {
				return 1
			}
			return
		}

		if !limitLow && !limitHigh {
			p := &memo[i][sum]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		lo := 0
		if limitLow {
			lo = int(num1[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(num2[i] - '0')
		}

		for d := lo; d <= hi; d++ {
			res = (res + dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)) % mod
		}
		return
	}
	return dfs(0, 0, true, true)
}
