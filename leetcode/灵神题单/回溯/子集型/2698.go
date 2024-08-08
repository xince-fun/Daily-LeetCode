package main

import "strconv"

var preSum [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n {
				return sum == i
			}
			x := 0
			for j := p; j < n; j++ {
				x = x*10 + int(s[j]-'0')
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		preSum[i] = preSum[i-1]
		if dfs(0, 0) {
			preSum[i] += i * i
		}
	}
}

func punishmentNumber(n int) int {
	return preSum[n]
}
