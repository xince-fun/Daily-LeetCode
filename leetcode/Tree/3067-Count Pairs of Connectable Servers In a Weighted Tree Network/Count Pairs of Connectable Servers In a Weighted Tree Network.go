package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) (ans []int) {
	type pair struct {
		y, weight int
	}
	gg := make([][]pair, len(edges))

	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		gg[x] = append(gg[x], pair{y, w})
		gg[y] = append(gg[y], pair{x, w})
	}

	ans = make([]int, len(edges))
	for i, g := range gg {
		cnt := 0
		var dfs func(int, int, int)

		dfs = func(x, fa, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, e := range gg[x] {
				if e.y != fa {
					dfs(e.y, x, sum+e.weight)
				}
			}
			return
		}

		sum := 0
		for _, e := range g {
			cnt = 0
			dfs(e.y, i, e.weight)
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}
