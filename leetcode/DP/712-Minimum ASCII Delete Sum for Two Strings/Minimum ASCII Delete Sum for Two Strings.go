package main

//			   dfs(i-1,j-1) s[i] == t[j]
// dfs(i,j) =  min(dfs(i-1, j)+int(s[i]), dfs(i, j-1)+int(t[i]))

func sum(str string) (ans int) {
	for _, s := range str {
		ans += int(s)
	}
	return
}

func minimumDeleteSum1(s, t string) int {
	n, m := len(s), len(t)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return sum(t[:j+1])
		}
		if j < 0 {
			return sum(s[:i+1])
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s[i] == s[j] {
			memo[i][j] = dfs(i-1, j-1)
		} else {
			memo[i][j] = min(dfs(i-1, j)+int(s[i]), dfs(i, j-1)+int(t[i]))
		}
		return memo[i][j]
	}
	return dfs(n-1, m-1)
}

func minimumDeleteSum2(s, t string) int {
	n, m := len(s), len(t)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 0; j <= m; j++ {
		f[0][j] = sum(t[:j])
	}
	for i, x := range s {
		f[i+1][0] = sum(s[:i+1])
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(f[i][j+1]+int(s[i]), f[i+1][j]+int(t[j]))
			}
		}
	}
	return f[n][m]
}

func minimumDeleteSum(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 0; j <= m; j++ {
		f[j] = sum(t[:j])
	}
	for i, x := range s {
		pre := f[0]
		f[0] = sum(s[:i+1])
		for j, y := range t {
			temp := f[j+1]
			if x == y {
				f[j+1] = pre
			} else {
				f[j+1] = min(f[j+1]+int(s[i]), f[j]+int(t[j]))
			}
			pre = temp
		}
	}
	return f[m]
}
