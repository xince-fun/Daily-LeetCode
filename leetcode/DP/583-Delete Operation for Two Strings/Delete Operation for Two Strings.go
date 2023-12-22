package main

//			= dfs(i-1,j-1) s[i] = t[j]
// dfs(i,j) = min(dfs(i-1,j),dfs(i,j-1))+1

func minDistance1(s, t string) int {
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
			memo[i][j] = min(dfs(i-1, j), dfs(i, j-1)) + 1
		}
		return memo[i][j]
	}
	return dfs(n-1, m-1)
}

func minDistance2(s, t string) int {
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
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}
	return f[n][m]
}

func minDistance3(s, t string) int {
	n, m := len(s), len(t)

	f := [2][]int{make([]int, m+1), make([]int, m+1)}
	for j := 0; j <= m; j++ {
		f[0][j] = j
	}
	for i, x := range s {
		f[(i+1)%2][0] = i + 1
		for j, y := range t {
			if x == y {
				f[(i+1)%2][j+1] = f[i%2][j]
			} else {
				f[(i+1)%2][j+1] = min(f[i%2][j+1], f[(i+1)%2][j]) + 1
			}
		}
	}
	return f[n%2][m]
}

func minDistance(s, t string) int {
	m := len(t)
	f := make([]int, m+1)
	for j := 0; j <= m; j++ {
		f[j] = j
	}
	for _, x := range s {
		pre := f[0]
		f[0]++
		for j, y := range t {
			temp := f[j+1]
			if x == y {
				f[j+1] = pre
			} else {
				f[j+1] = min(f[j+1], f[j]) + 1
			}
			pre = temp
		}
	}
	return f[m]
}
