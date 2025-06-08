package main

// 递归边界 dfs(-1, j) = 0, dfs(i, -1) = 0

func uniquePaths1(m int, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p
		}
		*p = dfs(i-1, j) + dfs(i, j-1)
		return *p
	}
	return dfs(m-1, n-1)
}

// f[i][0] f[0][j] 表示出界 f[1][0] = 1 // 首个位置

func uniquePaths2(m int, n int) int {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[1][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i+1][j+1] = f[i][j+1] + f[i+1][j]
		}
	}
	return f[m][n]
}

func uniquePaths3(m int, n int) int {
	f := [2][]int{make([]int, n+1), make([]int, n+1)}
	f[0][1] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[(i+1)%2][j+1] = f[i%2][j+1] + f[(i+1)%2][j]
		}
	}
	return f[m%2][n]
}

func uniquePaths(m int, n int) int {
	f := make([]int, n+1)
	f[1] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f[j+1] = f[j+1] + f[j]
		}
	}
	return f[n]
}
