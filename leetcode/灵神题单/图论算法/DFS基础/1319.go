package main

func makeConnected(n int, connections [][]int) int {
	g := make([][]int, n)
	for _, conn := range connections {
		x, y := conn[0], conn[1]
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
				size += dfs(x)
			}
		}
		return size
	}

	m := len(connections)
	cnt := 0
	total := 0
	for i, v := range vis {
		if !v {
			num := dfs(i)
			total += num - 1
			cnt++
		}
	}
	// cnt 个分区
	if m-total < cnt-1 {
		return -1
	}
	return cnt - 1
}
