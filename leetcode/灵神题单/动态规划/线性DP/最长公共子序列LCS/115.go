package main

// f[i][j] 表示以 i为结尾的s子序列，j为结尾的t子序列
// 当s[i] == t[j]时，可以匹配，也可以不匹配
// dfs(i,j) = dfs(i-1,j-1) + dfs(i-1,j)
// s[i] != t[j] 跳过s的当前位置
// dfs(i,j) = dfs(i-1,j)

func numDistinct1(s string, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		f[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] {
				f[i+1][j+1] = f[i][j] + f[i][j+1]
			} else {
				f[i+1][j+1] = f[i][j+1]
			}
		}
	}
	return f[n][m]
}

func numDistinct(s string, t string) int {
	const mod = 1_000_000_007
	n, m := len(s), len(t)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		f[0] = 1
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				f[j+1] = f[j] + f[j+1]
			}
		}
	}
	return f[m] % mod
}
