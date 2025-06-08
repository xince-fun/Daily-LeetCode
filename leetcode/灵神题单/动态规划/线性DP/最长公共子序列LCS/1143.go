package main

// 一般定义 f[i][j] 表示对 (s[:i],t[:j])的求解结果

// dfs(i,j) = dfs(i-1,j-1) + 1  s[i] = t[j]
//			= max(dfs(i-1,j), dfs(i,j-1)) s[i] != t[j]

func longestCommonSubsequence1(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// 表示其中一个字符串为空
		if i < 0 || j < 0 {
			return 0
		}
		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(n-1, m-1)
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// 表示其中一个字符串为空
		if i < 0 || j < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p > -1 {
			return *p
		}
		if text1[i] == text2[j] {
			*p = dfs(i-1, j-1) + 1
		} else {
			*p = max(dfs(i-1, j), dfs(i, j-1))
		}
		return *p
	}
	return dfs(n-1, m-1)
}

// f[i][j] = f[i-1][j-1] + 1  			s[i] == t[j]
//		   = max(f[i-1][j], f[i][j-1])	s[i] != t[j]

func longestCommonSubsequence3(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}

func longestCommonSubsequence4(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	f := [][]int{make([]int, m+1), make([]int, m+1)}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				f[(i+1)%2][j+1] = f[i%2][j] + 1
			} else {
				f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j])
			}
		}
	}
	return f[n%2][m]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		// 用一个 pre 存左上方向的值
		pre := f[0]
		for j := 0; j < m; j++ {
			tmp := f[j+1]
			if text1[i] == text2[j] {
				f[j+1] = pre + 1
			} else {
				f[j+1] = max(f[j+1], f[j])
			}
			pre = tmp
		}
	}
	return f[m]
}
