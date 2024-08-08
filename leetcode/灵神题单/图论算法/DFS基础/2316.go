package main

func countPairs(n int, edges [][]int) (ans int64) {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)

	var dfs func(int) int
	dfs = func(x int) int {
		size := 1
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] {
				size += dfs(y)
			}
		}
		return size
	}
	cnt := int64(0)
	for i, v := range vis {
		if !v {
			size := dfs(i)
			ans += int64(size) * cnt
			cnt += int64(size)
		}
	}
	return
}
