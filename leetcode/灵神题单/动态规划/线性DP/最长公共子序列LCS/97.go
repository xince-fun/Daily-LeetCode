package main

func isInterleave1(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	memo := make([][]*bool, m+1)
	for i := range memo {
		memo[i] = make([]*bool, n+1)
	}

	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return s2[:j+1] == s3[:j+1]
		}
		if j < 0 {
			return s1[:i+1] == s3[:i+1]
		}
		if memo[i][j] != nil {
			return *memo[i][j]
		}
		res := (dfs(i-1, j) && s1[i] == s3[i+j+1] || (dfs(i, j-1) && s2[j] == s3[i+j+1]))
		memo[i][j] = &res
		return res
	}
	return dfs(m-1, n-1)
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i < n; i++ {
		f[i+1][0] = f[i][0] && s1[i] == s3[i]
	}
	for j := 0; j < m; j++ {
		f[0][j+1] = f[0][j] && s2[j] == s3[j]
	}
	for i, x := range s1 {
		for j, y := range s2 {
			f[i+1][j+1] = (f[i][j+1] && byte(x) == s3[i+j+1]) || (f[i+1][j] && byte(y) == s3[i+j+1])
		}
	}
	return f[n][m]
}
