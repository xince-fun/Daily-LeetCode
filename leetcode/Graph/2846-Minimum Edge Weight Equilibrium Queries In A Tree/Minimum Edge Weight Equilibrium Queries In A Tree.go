package main

import "slices"

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}

	const mx = 14
	type pair struct {
		p   int
		cnt [26]int
	}
	pa := make([][mx]pair, n)
	depth := make([]int, n)
	var build func(int, int, int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		depth[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].cnt[e.wt] = 1
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][v+1].p = pp.p
				for j := 0; j < 26; j++ {
					pa[v][i+1].cnt[j] = p.cnt[j] + pp.cnt[j]
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	f := func(v, w int) int {
		pathLen := depth[v] + depth[w]
		cnt := [26]int{}
		if depth[v] > depth[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (depth[w]-depth[v])>>i&1 > 0 {
				p := pa[w][i]
				for j := 0; j < 26; j++ {
					cnt[j] += p.cnt[j]
				}
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					for j := 0; j < 26; j++ {
						cnt[j] += pv.cnt[j] + pw.cnt[j]
					}
					v, w = pv.p, pw.p
				}
			}
			for j := 0; j < 26; j++ {
				cnt[j] += pa[v][0].cnt[j] + pa[w][0].cnt[j]
			}
			v = pa[v][0].p
		}
		return pathLen - depth[v]*2 - slices.Max(cnt[:])
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q[0], q[1])
	}
	return ans
}
