package main

import "strconv"

// 选不选逗号的问题
func isAdditiveNumber(num string) (ans bool) {
	n := len(num)

	path := make([]int, 0)
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			if len(path) >= 3 {
				ans = true
			}
			return
		}
		if i < n-1 {
			dfs(i+1, start)
		}

		s := num[start : i+1]
		str, _ := strconv.Atoi(s)
		if len(path) >= 2 {
			if path[len(path)-1]+path[len(path)-2] != str {
				return
			}
		}
		path = append(path, str)
		dfs(i+1, i+1)
		path = path[:len(path)-1]
	}

	dfs(0, 0)
	return
}
