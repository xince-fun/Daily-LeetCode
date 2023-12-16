package main

// 从2n个位置中选n个位置放左括号
// 选 左括号
// 不选 右括号

// 1. 当前操作？枚举当前path[i]是左括号还是右括号
// 2. 子问题？构造字符串>=i的部分
// 3. 下一个子问题？构造字符串>=i+1的部分

// 还需要记录左括号的个数 open，从而方便判断
// 需要选n个左括号
// 只要open < n 就可以选左括号

// 右括号个数为 i - open
// 如果右括号的个数小于左括号的个数
// i < open < open， 就可以选左扩号

func generateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		// 还可以加左括号
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
