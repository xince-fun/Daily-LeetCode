package main

// 选不选的角度
func letterCasePermutation1(s string) (ans []string) {
	n := len(s)

	path := []byte(s)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		dfs(i + 1)
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			path[i] ^= 32
			dfs(i + 1)
			path[i] ^= 32
		}
	}
	dfs(0)
	return
}

// 从答案的角度
func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	path := []byte(s)
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, string(path))
		for j := i; j < n; j++ {
			if (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z') {
				path[j] ^= 32
				dfs(j + 1)
				path[j] ^= 32
			}
		}
	}
	dfs(0)
	return
}
