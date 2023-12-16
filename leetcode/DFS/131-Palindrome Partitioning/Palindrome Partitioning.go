package main

import "fmt"

func partition1(s string) (ans [][]string) {
	n := len(s)
	path := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
		}
		for j := i; j < n; j++ {
			t := s[i : j+1]
			if isParli(t) {
				path = append(path, t)
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

func isParli(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

// s = "aab"
//

func partition(s string) (ans [][]string) {
	n := len(s)
	path := []string{}

	var dfs func(int, int)
	// start 表示这段字段开始的位置
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		// 不选 i和i+1之间的逗号
		if i < n-1 {
			dfs(i+1, start)
		}

		// 把 s[i] 作为子串的最后一个字符
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
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
