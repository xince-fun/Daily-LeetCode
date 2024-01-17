package main

import "strconv"

func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool) int
	f = func(i, cnt1 int, isLimit bool) (res int) {
		if i == m {
			return cnt1
		}
		if !isLimit {
			dv := &dp[i][cnt1]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for d := 0; d <= up; d++ { // 枚举要填入的数字 d
			c := cnt1
			if d == 1 {
				c++
			}
			res += f(i+1, c, isLimit && d == up)
		}
		return
	}
	return f(0, 0, true)
}
