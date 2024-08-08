package main

// dp with bitmasking
// last_ind  / the mask of numbers

func specialPerm1(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	memo := make([][]int, u)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(s, i int) (res int) {
		if s == 0 {
			return 1
		}
		p := &memo[s][i]
		if *p != -1 {
			return *p
		}
		for j, x := range nums {
			if s>>j&1 > 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
				res += dfs(s^(1<<j), j)
			}
		}
		*p = res
		return res
	}
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % 1_000_000_007
}

func specialPerm(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	f := make([][]int, u)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range nums {
		f[0][i] = 1
	}
	for s := 1; s < u; s++ {
		for i, pre := range nums {
			if s>>i&1 != 0 {
				continue
			}
			for j, x := range nums {
				if s>>j&1 != 0 && (pre%x == 0 || x%pre == 0) {
					f[s][i] += f[s^(1<<j)][j]
				}
			}
		}
	}
	for i := range nums {
		ans += f[u^(1<<i)][i]
	}
	return ans % 1_000_000_007

}
