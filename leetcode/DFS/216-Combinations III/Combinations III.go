package main

func combinationSum31(k int, n int) (ans [][]int) {
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path)
		if i < d {
			return
		}
		if len(path) == k && sum(path) == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j := i; j > 0; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(9)
	return
}

func sum(path []int) (ans int) {
	for _, p := range path {
		ans += p
	}
	return
}

func combinationSum3(k int, n int) (ans [][]int) {
	path := []int{}

	var dfs func(int, int)
	dfs = func(i, t int) {
		d := k - len(path)
		if t < 0 || t > (i*2-d+1)*d/2 {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j > d-1; j-- {
			path = append(path, j)
			dfs(j-1, t-j)
			path = path[:len(path)-1]
		}
	}
	dfs(9, n)
	return
}
