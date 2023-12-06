package main

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// bob 的移动路径是确定的
	// DFS 求出 bob 路径，计算到达每个点的时间，不在路径的初始化为无穷或者n
	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = n
	}
	var dfs func(int, int, int) bool // 返回是否找到 0
	dfs = func(x, fa, t int) bool {
		if x == 0 {
			bobTime[x] = t
			return true
		}
		found0 := false
		for _, y := range g[x] {
			if y != fa {
				if dfs(y, x, t+1) {
					found0 = true
				}
			}
		}
		if found0 {
			bobTime[x] = t
		}
		return found0
	}
	dfs(bob, -1, 0)

	// 小技巧，给 根节点加一个节点
	g[0] = append(g[0], -1)
	ans := math.MinInt32

	var dfs_alice func(int, int, int, int)
	dfs_alice = func(x, fa, aliceTime, total int) {
		if aliceTime < bobTime[x] {
			total += amount[x]
		} else if aliceTime == bobTime[x] {
			total += amount[x] / 2
		}
		if len(g[x]) == 1 {
			ans = max(ans, total)
			return
		}
		for _, y := range g[x] {
			if y != fa {
				dfs_alice(y, x, aliceTime+1, total)
			}
		}
	}
	dfs_alice(0, -1, 0, 0)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
