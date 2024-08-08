package main

import "sort"

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) int {
	n := len(grid)

	return sort.Search(n*n, func(t int) bool {
		if t < grid[0][0] {
			return false
		}
		g := make([][]bool, n)
		for i := 0; i < n; i++ {
			g[i] = make([]bool, n)
			for j := 0; j < n; j++ {
				g[i][j] = t >= grid[i][j]
			}
		}
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		visited[0][0] = true
		type pair struct{ x, y int }
		q := []pair{}
		q = append(q, pair{0, 0})

		for len(q) > 0 {
			e := q[0]
			q = q[1:]

			if e.x == n-1 && e.y == n-1 {
				return true
			}

			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if 0 <= nx && nx < n && 0 <= ny && ny < n && !visited[nx][ny] && g[nx][ny] {
					q = append(q, pair{nx, ny})
					visited[nx][ny] = true
				}
			}
		}
		return false
	})
}
