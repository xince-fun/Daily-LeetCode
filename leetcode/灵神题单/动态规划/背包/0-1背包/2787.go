package main

// n 表示成一些 n1^x + n2^x + n3^x + n4^x
// target = 10
// 从i个数中，选一些数，它们的幂和为target

func numberOfWays1(n int, x int) int {
	const mod = 1_000_000_007
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 1 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		num := pow(i, x)
		if c < num {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = dfs(i-1, c) + dfs(i-1, c-num)
		}
		return memo[i][c]
	}
	return dfs(n, n) % mod
}

func numberOfWays2(n int, x int) int {
	const mod = 1_000_000_007
	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[1][0] = 1

	for i := 1; i <= n; i++ {
		num := pow(i, x)
		for c := 0; c <= n; c++ {

			if c < num {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i][c-num]
			}
		}
	}
	return f[n+1][n] % mod
}

func numberOfWays(n int, x int) int {
	const mod = 1_000_000_007
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		num := pow(i, x)
		for c := n; c >= num; c-- {
			f[c] = f[c] + f[c-num]
		}
	}
	return f[n] % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
