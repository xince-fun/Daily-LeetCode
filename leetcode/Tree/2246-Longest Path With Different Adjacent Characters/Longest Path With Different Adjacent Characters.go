package main

func longestPath(parent []int, s string) int {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}
	ans := 0
	var dfs func(int) int
	dfs = func(x int) int {
		xLen := 0
		for _, y := range g[x] {
			yLen := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, xLen+yLen)
				xLen = max(xLen, yLen)
			}
		}
		return xLen
	}
	dfs(0)
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
