package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n) // g 表示 x 的所有邻居
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	size := make([]int, n)
	var dfs func(int, int, int)

	dfs = func(x, fa, depth int) {
		ans[0] += depth // depth 为0到x的距离
		size[x] = 1
		for _, y := range g[x] { // x的邻居y
			if y != fa {
				dfs(y, x, depth+1) // x 是 y的父节点
				size[x] += size[y] // 累加 x 的儿子 y的子树大小
			}
		}
	}
	dfs(0, -1, 0)

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				ans[y] = ans[x] + n - 2*size[y]
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
