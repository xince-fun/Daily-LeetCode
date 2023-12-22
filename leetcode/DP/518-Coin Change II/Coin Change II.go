package main

// 记忆化搜索
func change1(amount int, coins []int) int {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if memo[i][c] != -1 {
			return memo[i][c]
		}
		if c < coins[i] {
			memo[i][c] = dfs(i-1, c)
		} else {
			memo[i][c] = dfs(i-1, c) + dfs(i, c-coins[i])
		}
		return memo[i][c]
	}
	return dfs(n-1, amount)
}

func change2(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i+1][c-coin]
			}
		}
	}
	return f[n][amount]
}

func change3(amount int, coins []int) int {
	n := len(coins)
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	f[0][0] = 1
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] + f[(i+1)%2][c-coin]
			}
		}
	}
	return f[n%2][amount]
}

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c >= coin {
				f[c] = f[c] + f[c-coin]
			}
		}
	}
	return f[amount]
}
