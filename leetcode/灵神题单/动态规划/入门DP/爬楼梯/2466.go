package main

// 不同好字符串
// 最后次在末尾添加zero个零，则 i-zero 的个数， i-zero的个数
// 最后次在末尾添加one个一，则 i-one 的个数， i-one的个数

// dfs(i) = dfs(i-zero) + dfs(i-one)

// dfs(0) 为 1，表示当长度为0的时候，构造的方法就只有一种。

// 当前操作？

func countGoodStrings1(low int, high int, zero int, one int) (ans int) {
	const mod = 1_000_000_007
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i <= 0 {
			return 1
		}
		if i >= zero {
			res += dfs(i-zero) % mod
		}
		if i >= one {
			res += dfs(i-one) % mod
		}
		if i >= low {
			ans += res % mod
		}
		return res
	}
	dfs(high)
	return ans
}

func countGoodStrings2(low int, high int, zero int, one int) (ans int) {
	const mod = 1_000_000_007
	memo := make([]int, high+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i <= 0 {
			return 1
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		if i >= zero {
			res = (res + dfs(i-zero)) % mod
		}
		if i >= one {
			res = (res + dfs(i-one)) % mod
		}
		*p = res
		return res % mod
	}
	for i := low; i <= high; i++ {
		ans = (ans + dfs(i)) % mod
	}
	return
}

func countGoodStrings(low int, high int, zero int, one int) (ans int) {
	const mod = 1_000_000_007
	f := make([]int, high+1)
	f[0] = 1
	for i := 1; i <= high; i++ {
		if i >= one {
			f[i] = (f[i] + f[i-one]) % mod
		}
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % mod
		}
		if i >= low {
			ans = (ans + f[i]) % mod
		}
	}
	return ans
}
