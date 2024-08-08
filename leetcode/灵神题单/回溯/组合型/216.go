package main

func combinationSum3(k int, n int) (ans [][]int) {
	path := []int{}

	var dfs func(int, int)
	dfs = func(i, sum int) {
		d := k - len(path)
		if sum < 0 || sum > (i*2-d+1)*d/2 {
			return
		}
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j > d-1; j-- {
			path = append(path, j)
			dfs(j-1, sum-j)
			path = path[:len(path)-1]
		}
	}
	dfs(9, n)
	return
}
