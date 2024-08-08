package main

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)
	// 以圆心 找半径，覆盖则加入
	for i, b := range bombs {
		x1, y1, r := b[0], b[1], b[2]
		for j, b2 := range bombs {
			if i == j {
				continue
			}
			x2, y2 := b2[0], b2[1]
			if distance(x1, y1, x2, y2) <= r*r {
				g[i] = append(g[i], j)
			}
		}
	}
	var size int
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		size++
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	ans := 0
	for i := range g {
		clear(vis)
		size = 0
		dfs(i)
		ans = max(ans, size)
	}

	return ans
}

func distance(x1, y1, x2, y2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}
