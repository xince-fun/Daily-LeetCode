package main

// 回溯三问
// 1. 当前操作？枚举path[i]要填入的字母
// 2. 子问题？ 构造字符串>=i的部分
// 3. 下一个子问题？ 构造字符串>=i+1的部分
// dfs(i) -> dfs(i+1)

var mapping = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	ans := []string{}

	path := make([]rune, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range mapping[digits[i]-'0'] {
			path[i] = c
			dfs(i + 1)
		}
	}
	dfs(0)
	return ans
}
