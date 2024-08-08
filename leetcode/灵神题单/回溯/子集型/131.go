package main

func partition1(s string) (ans [][]string) {
	n := len(s)
	path := []string{}

	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		// 不选
		if i < n-1 {
			dfs(i+1, start)
		}

		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}

func partition(s string) (ans [][]string) {
	n := len(s)
	path := []string{}

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		for j := i; j < n; j++ {
			if isPalin(s[i : j+1]) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func isPalin(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
