package main

// 树形dp?
func minimumFuelCost(roads [][]int, seats int) int64 {
	var ans int64
	n := len(roads) + 1
	g := make([][]int, n) // g 表示 x 的所有邻居
	for _, r := range roads {
		x, y := r[0], r[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int

	// depth 为 0到x的距离
	dfs = func(x, fa int) int {
		size := 1
		for _, y := range g[x] {
			if y != fa {
				size += dfs(y, x)
			}
		}
		if x > 0 {
			ans += int64((size + seats - 1) / seats)
		}
		return size
	}
	dfs(0, -1)

	return ans
}
