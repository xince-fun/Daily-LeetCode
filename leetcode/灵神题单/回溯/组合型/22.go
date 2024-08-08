package main

// 从2n个位置种选n个位置放左括号
func generateParenthesis(n int) (ans []string) {
	m := 2 * n
	path := make([]byte, m)

	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open {
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}
