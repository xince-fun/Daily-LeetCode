package main

import (
	"fmt"
	"strconv"
)

// 选不选逗号的问题

func punishmentNumber1(n int) (ans int) {
	var str string
	var m, count, sq, num int
	var vis [1001]bool
	var dfs func(int, int)

	dfs = func(i, start int) {
		if i == m {
			if count == num && !vis[num] {
				ans += sq
				vis[num] = true
			}
			return
		}
		if i < m-1 {
			dfs(i+1, start)
		}
		num := str[start : i+1]
		str, _ := strconv.Atoi(num)
		if num != strconv.Itoa(str) {
			return
		}
		count += str
		dfs(i+1, i+1)
		count -= str
	}

	for i := 1; i <= n; i++ {
		sq, count, num = i*i, 0, i
		str = strconv.Itoa(sq)
		m = len(str)
		dfs(0, 0)
	}
	return
}

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

func punishmentNumber(n int) (ans int) {
	return preSum[n]
}

func main() {
	n := 91
	fmt.Println(punishmentNumber(n))
}
