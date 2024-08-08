package main

func countCompleteComponents1(n int, edges [][]int) (ans int) {
	vis := make([]bool, n)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int) (int, int)
	dfs = func(x int) (int, int) {
		size := 1
		cnt := 0
		vis[x] = true
		for _, y := range g[x] {
			cnt += 1
			if !vis[y] {
				s, c := dfs(y)
				size += s
				cnt += c
			}
		}
		return size, cnt
	}

	for i, v := range vis {
		if !v {
			s, c := dfs(i)
			if c == s*(s-1) {
				ans++
			}
		}
	}
	return
}

func countCompleteComponents(n int, edges [][]int) (ans int) {
	vis := make([]bool, n)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	v, e := 0, 0
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		v++
		e += len(g[x])
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}
	for i, b := range vis {
		if !b {
			v, e = 0, 0
			dfs(i)
			if e == v*(v-1) {
				ans++
			}
		}
	}

	return
}
