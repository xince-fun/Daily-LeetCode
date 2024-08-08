package main

func removeInvalidParentheses(s string) (ans []string) {
	maxl := 0
	path := []byte{}
	m := map[string]bool{}
	var dfs func(int, int, int)
	dfs = func(i, lopen, ropen int) {
		if i == len(s) {
			if lopen == ropen {
				if len(path) >= maxl {
					maxl = len(path)
					m[string(path)] = true
				}
			}
			return
		}
		if 'a' <= s[i] && s[i] <= 'z' {
			path = append(path, s[i])
			dfs(i+1, lopen, ropen)
			path = path[:len(path)-1]
		}
		if s[i] == '(' {
			path = append(path, s[i])
			dfs(i+1, lopen+1, ropen)
			path = path[:len(path)-1]
			dfs(i+1, lopen, ropen)
		}
		if s[i] == ')' {
			if lopen > ropen {
				path = append(path, s[i])
				dfs(i+1, lopen, ropen+1)
				path = path[:len(path)-1]
				dfs(i+1, lopen, ropen)
			} else {
				dfs(i+1, lopen, ropen)
			}
		}
	}
	dfs(0, 0, 0)
	for k := range m {
		ans = append(ans, k)
	}
	return
}
