package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	g[1] = []int{0} // 减少特判
	var ans float64
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int, int, int) bool
	dfs = func(x, fa, lefT, prod int) bool {
		// 此时到达target 时间为0 或者 叶子节点
		if x == target && (lefT == 0 || len(g[x]) == 1) {
			ans = 1 / float64(prod)
			return true
		}
		if x == target || lefT == 0 {
			return false
		}
		for _, y := range g[x] {
			if y != fa && dfs(y, x, lefT-1, prod*(len(g[x])-1)) {
				return true // 找到 target
			}
		}
		return false // 没找到 target
	}

	dfs(0, -1, t, 1)
	return ans
}
