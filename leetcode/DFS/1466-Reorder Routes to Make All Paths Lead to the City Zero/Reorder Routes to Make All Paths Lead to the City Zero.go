package main

func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n)
	//  对于 a,b 视为 b,a
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], [2]int{c[1], 1})
		g[c[1]] = append(g[c[1]], [2]int{c[0], 0})
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) (ans int) {
		for _, y := range g[x] {
			if b, c := y[0], y[1]; b != fa {
				ans += c + dfs(b, x)
			}
		}
		return
	}
	return dfs(0, -1)
}
