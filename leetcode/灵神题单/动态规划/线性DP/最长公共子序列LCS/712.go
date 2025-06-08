package main

//

func minimumDeleteSum1(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	sum := func(s string) int {
		sum := 0
		for _, c := range s {
			sum += int(c)
		}
		return sum
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return sum(s2[:j+1])
		}
		if j < 0 {
			return sum(s1[:i+1])
		}
		p := &memo[i][j]
		if *p > -1 {
			return *p
		}
		if s1[i] == s2[j] {
			*p = dfs(i-1, j-1)
		} else {
			*p = min(dfs(i-1, j)+int(s1[i]), dfs(i, j-1)+int(s2[j]))
		}
		return *p
	}
	return dfs(n-1, m-1)
}

func minimumDeleteSum2(s1 string, s2 string) int {
	n, m := len(s1), len(s2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = f[0][j-1] + int(s2[j-1])
	}
	for i := 0; i < n; i++ {
		f[i+1][0] = f[i][0] + int(s1[i])
		for j := 0; j < m; j++ {
			if s1[i] == s2[j] {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1]+int(s1[i]), f[i+1][j]+int(s2[j]))
			}
		}
	}
	return f[n][m]
}

func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)

	f := make([]int, m+1)
	// i=0的j
	for j := 1; j <= m; j++ {
		f[j] = f[j-1] + int(s2[j-1])
	}
	sum := 0
	for i := 0; i < n; i++ {
		// pre 是 f[i][j]
		pre := f[0]
		f[0] = sum + int(s1[i])
		sum = f[0]
		for j := 0; j < m; j++ {
			tmp := f[j+1]
			if s1[i] == s2[j] {
				f[j+1] = pre
			} else {
				f[j+1] = min(f[j+1]+int(s1[i]), f[j]+int(s2[j]))
			}
			pre = tmp
		}
	}
	return f[m]
}
