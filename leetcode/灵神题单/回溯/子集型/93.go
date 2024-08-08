package main

import (
	"strconv"
)

func restoreIpAddresses(s string) (ans []string) {
	n := len(s)

	isValid := func(s string) bool {
		num, _ := strconv.Atoi(s)
		if s != strconv.Itoa(num) {
			return false
		}
		if 0 <= num && num <= 255 {
			return true
		}
		return false
	}

	path := ""
	cnt := 0
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			if cnt == 4 {
				ans = append(ans, path[:len(path)-1])
			}
			return
		}
		if i < n-1 {
			dfs(i+1, start)
		}

		str := s[start : i+1]
		if isValid(str) {
			path += str + "."
			cnt++
			dfs(i+1, i+1)
			cnt--
			length := len(str + ".")
			path = path[:len(path)-length]
		}
	}

	dfs(0, 0)
	return ans
}
