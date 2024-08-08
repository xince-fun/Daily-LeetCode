package main

func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for d := 0; q != nil; d++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v > 0 {
				p := patience[v]
				t := p
				if 2*d%p > 0 {
					t = 2 * d % p
				}
				ans = max(ans, 4*d-t+1)
			}
			for _, w := range g[v] {
				if !vis[w] {
					vis[w] = true
					q = append(q, w)
				}
			}
		}
	}
	return
}
