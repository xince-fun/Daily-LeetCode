package main

// 区间 DP

// 选或不选 从两侧向内缩小问题规模
// 枚举选哪个 分割成多个规模更小的子问题

// 定义dfs(i,j)表示从s[i]到s[j]的最长回文子序列的长度

//				dfs(i+1,j-1) +2  s[i] = s[j]
// dfs((i,j) =  max(dfs(i+1,j),dfs(i,j-1)) s[i] != s[j]

// 递归边界 dfs(i,i) = 1
// dfs(i+1, i) = 0

func longestPalindromeSubseq1(s string) int {
	n := len(s)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if s[i] == s[j] {
			return dfs(i+1, j-1) + 2
		}
		return max(dfs(i+1, j), dfs(i, j-1))
	}
	return dfs(0, n-1)
}

func longestPalindromeSubseq(s string) int {
	n := len(s)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i > 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
			}
		}
	}
	return f[0][n-1]
}
