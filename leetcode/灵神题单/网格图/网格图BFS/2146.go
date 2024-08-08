package main

import "slices"

var dirs = []struct{ x, y int }{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	type pair struct{ x, y, t int }
	m, n := len(grid), len(grid[0])

	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}

	q := []pair{}
	ans := []pair{}
	q = append(q, pair{start[0], start[1], 0})

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		if dis[e.x][e.y] != -1 {
			continue
		}
		dis[e.x][e.y] = e.t
		if pricing[0] <= grid[e.x][e.y] && grid[e.x][e.y] <= pricing[1] {
			ans = append(ans, pair{e.x, e.y, e.t})
		}
		for _, dir := range dirs {
			nx, ny := e.x+dir.x, e.y+dir.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] >= 1 && dis[nx][ny] == -1 {
				q = append(q, pair{nx, ny, e.t + 1})
			}
		}
	}
	slices.SortFunc(ans, func(a, b pair) int {
		if a.t != b.t {
			return a.t - b.t
		}
		if grid[a.x][a.y] != grid[b.x][b.y] {
			return grid[a.x][a.y] - grid[b.x][b.y]
		}
		if a.x != b.x {
			return a.x - b.x
		}
		return a.y - b.y
	})
	a := [][]int{}
	for i := 0; i < min(k, len(ans)); i++ {
		a = append(a, []int{ans[i].x, ans[i].y})
	}
	return a
}
