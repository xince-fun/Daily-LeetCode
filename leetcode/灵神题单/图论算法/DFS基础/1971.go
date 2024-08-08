package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	var dfs func(int) bool
	dfs = func(x int) bool {
		if x == destination {
			return true
		}
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] && dfs(y) {
				return true
			}
		}
		return false
	}
	return dfs(source)
}
