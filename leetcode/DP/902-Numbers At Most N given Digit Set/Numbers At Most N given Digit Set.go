package main

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = -1
	}
	// 返回从 i 开始填数字，i 前面填的数字的集合是 mask，能构造出的特殊整数的树木
	// isLimit 表示前面填的数字是否都是 n 对应位上的， 如果为true， 那么当前位至多为 int(s[i])，否则至多为 '9'
	// isNum 表示前面是否填了数字（是否跳过），如果为 true， 那么当前位可以从 0开始，如果为 false，那么我们可以跳过，或者从 1开始填数字
	var f func(int, bool, bool) int
	f = func(i int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum {
			res += f(i+1, false, false)
		}

		up := byte('9')
		if isLimit {
			up = s[i]
		}
		for _, d := range digits {
			if d[0] > up {
				break
			}
			res += f(i+1, isLimit && d[0] == up, true)
		}
		return
	}
	return f(0, true, false)
}
