package main

//			= dfs(i-1,j-1) s[i] == t[j]
// dfs(i,j) = min(dfs(i,j-1),dfs(i-1,j),dfs(i-1,j-1)) + 1// s[i] != t[j]
//                    插入        删除        替换
//				   跟删另一个一样

func minDistance1(s string, t string) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
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
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s[i] == t[j] {
			memo[i][j] = dfs(i-1, j-1)
		} else {
			memo[i][j] = min(min(dfs(i, j-1), dfs(i-1, j)), dfs(i-1, j-1)) + 1
		}
		return memo[i][j]
	}
	return dfs(n-1, m-1)
}

func minDistance(s string, t string) int {
	n, m := len(s), len(t)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)

	}
	for j := 0; j <= m; j++ {
		f[0][j] = j
	}

	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(min(f[i+1][j], f[i][j+1]), f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}
