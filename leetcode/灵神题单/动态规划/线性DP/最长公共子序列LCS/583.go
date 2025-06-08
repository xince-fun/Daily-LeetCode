package main

// dfs(i,j) = dfs(i-1,j-1)  s[i] == t[j]
//			= min(dfs(i-1,j),dfs(i,j-1)) + 1 s[i] != t[j]
//				  删s			删t

func minDistance1(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
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
			*p = min(dfs(i-1, j), dfs(i, j-1)) + 1
		}
		return *p
	}
	return dfs(n-1, m-1)
}

func minDistance2(word1 string, word2 string) int {
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
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}
	return f[n][m]
}

func minDistance(word1, word2 string) int {
	n, m := len(word1), len(word2)

	f := make([]int, m+1)
	for j := 0; j <= m; j++ {
		f[j] = j
	}
	for i := 0; i < n; i++ {
		f[0] = i + 1
		pre := i
		for j := 0; j < m; j++ {
			tmp := f[j+1]
			if word1[i] == word2[j] {
				f[j+1] = pre
			} else {
				f[j+1] = min(f[j+1], f[j]) + 1
			}
			pre = tmp
		}
	}
	return f[m]
}
