package main

// s = horse
// t = ros

// dfs(i,j) = dfs(i-1,j-1)
//			= min(dfs(i, j-1), dfs(i-1,j), dfs(i-1,j-1)) + 1
//						插入        删除         替换

func minDistance1(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		// 把text2删了
		if i < 0 {
			return j + 1
		}
		// 把text1删了
		if j < 0 {
			return i + 1
		}
		p := &memo[i][j]
		if *p > -1 {
			return *p
		}
		if word1[i] == word2[j] {
			*p = dfs(i-1, j-1)
		} else {
			*p = min(dfs(i, j-1), dfs(i-1, j), dfs(i-1, j-1)) + 1
		}
		return *p
	}
	return dfs(n-1, m-1)
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		f[0][j] = j
	}
	for i := 0; i < n; i++ {
		f[i+1][0] = i + 1
		for j := 0; j < m; j++ {
			if word1[i] == word2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i+1][j], f[i][j+1], f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}
