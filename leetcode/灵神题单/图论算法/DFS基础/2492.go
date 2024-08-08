package main

import "math"

func minScore(n int, roads [][]int) (ans int) {
	ans = math.MaxInt
	type pair struct{ y, d int }
	g := make([][]pair, n+1)
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x] = append(g[x], pair{y, d})
		g[y] = append(g[y], pair{x, d})
	}

	vis := make([]bool, n+1)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, y := range g[x] {
			ans = min(ans, y.d)
			if !vis[y.y] {
				dfs(y.y)
			}
		}
	}

	dfs(1)
	return
}
