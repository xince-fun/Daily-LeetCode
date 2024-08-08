package main

// 正向
func getAncestors1(n int, edges [][]int) [][]int {
	ans := make([][]int, n)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
	}

	vis := make([]int, n)
	start := 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = start + 1
		for _, y := range g[x] {
			if vis[y] != start+1 {
				ans[y] = append(ans[y], start)
				dfs(y)
			}
		}
	}

	for ; start < n; start++ {
		dfs(start)
	}
	return ans
}

// 反向
func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	ans := make([][]int, n)
	for i := range ans {
		clear(vis)
		dfs(i)
		vis[i] = false
		for j, b := range vis {
			if b {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}
