package main

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges))
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses {
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 {
					cnt0++
				}
				dfs(y, x)
			}
		}
	}

	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}
