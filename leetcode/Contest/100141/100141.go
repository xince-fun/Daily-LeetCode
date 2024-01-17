package main

import "sort"

/*
给你m个数，从中选3个数，这三个数之积最大，如果最大值是负数就返回0.

1. 0个负数 3个最大的正数
2. 2个最小的负数 1个最大的正数
3. [0]

三个取最大值
*/

func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)

	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int64, n)
	for i := range ans {
		ans[i] = 1
	}

	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		a := []int{cost[x]}
		for _, y := range g[x] {
			if y != fa {
				a = append(a, dfs(y, x))
			}
		}
		sort.Ints(a)
		m := len(a)
		if m >= 3 {
			ans[x] = int64(max(max(a[m-1]*a[m-2]*a[m-3], a[0]*a[1]*a[m-1]), 0))
		}
		if m > 5 {
			a = append(a[:2], a[m-3:]...)
		}
		return a
	}
	dfs(0, -1)
}
