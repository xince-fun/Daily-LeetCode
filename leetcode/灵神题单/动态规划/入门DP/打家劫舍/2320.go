package main

// 乘法原理，两边互不干扰，统计一边 相乘即是答案

// 两种情况：
// 这里放房子： 前i-2得到的方法数
// 这里不放房子： 前i-1得到的方法数

// dfs(0)=1，表示没有地块的时候，只有都不放一种方案

func countHousePlacements1(n int) int {

	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 0 {
			return 1
		}
		res := dfs(i-1) + dfs(i-2)
		return res
	}
	a := dfs(n)
	return a * a
}

func countHousePlacements2(n int) int {
	const mod = 1_000_000_007
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 0 {
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		res := (dfs(i-1) + dfs(i-2)) % mod
		*p = res
		return res
	}
	a := dfs(n)
	return a * a % mod
}

// f[i] = f[i-1]+f[i-2]
// f[i+2] = f[i+1] + f[i]

func countHousePlacements3(n int) int {
	const mod = 1_000_000_007
	f := make([]int, n+1)
	f[0], f[1] = 1, 2
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}
	return f[n] * f[n] % mod
}

func countHousePlacements(n int) int {
	const mod = 1_000_000_007
	f0, f1 := 1, 2
	for i := 2; i <= n; i++ {
		f0, f1 = f1, (f0+f1)%mod
	}
	return f1 * f1 % mod
}
