package main

import "math"

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct{ to, color int }
	g := make([][]pair, n)
	for _, e := range redEdges {
		x, y := e[0], e[1]
		g[x] = append(g[x], pair{y, 0})
	}
	for _, e := range blueEdges {
		x, y := e[0], e[1]
		g[x] = append(g[x], pair{y, 1})
	}

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = math.MaxInt
	}
	cnt := 0
	q := []pair{{0, -1}}
	vis := make([][2]bool, n)
	for len(q) > 0 {
		cnt++
		tmp := q
		q = []pair{}
		for _, pp := range tmp {
			for _, p := range g[pp.to] {
				if p.color != pp.color && !vis[p.to][p.color] {
					vis[p.to][p.color] = true
					q = append(q, p)
					ans[p.to] = min(ans[p.to], cnt)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if ans[i] == math.MaxInt {
			ans[i] = -1
		}
	}
	return ans
}
