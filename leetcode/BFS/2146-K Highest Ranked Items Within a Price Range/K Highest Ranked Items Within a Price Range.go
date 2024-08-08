package main

import "sort"

var dirs = []struct{ x, y int }{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) (ans [][]int) {
	type pair struct{ x, y int }
	m, n := len(grid), len(grid[0])

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	type r struct{ x, y, dis, price int }
	res := []r{}

	if grid[start[0]][start[1]] >= pricing[0] && grid[start[0]][start[1]] <= pricing[1] {
		res = append(res, r{start[0], start[1], 0, grid[start[0]][start[1]]})
	}

	q := []pair{}
	q = append(q, pair{start[0], start[1]})
	vis[start[0]][start[1]] = true

	count := 0
	for len(q) > 0 {
		length := len(q)
		count++
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != 0 && !vis[nx][ny] {
					vis[nx][ny] = true
					q = append(q, pair{nx, ny})
					if grid[nx][ny] >= pricing[0] && grid[nx][ny] <= pricing[1] {
						res = append(res, r{nx, ny, count, grid[nx][ny]})
					}
				}
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].dis != res[j].dis {
			return res[i].dis < res[j].dis
		}
		if res[i].price != res[j].price {
			return res[i].price < res[j].price
		}
		if res[i].x != res[j].x {
			return res[i].x < res[j].x
		}
		return res[i].y < res[j].y
	})
	for i := 0; i < min(k, len(res)); i++ {
		ans = append(ans, []int{res[i].x, res[i].y})
	}
	return
}
