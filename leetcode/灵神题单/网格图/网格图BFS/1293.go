package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, y, count int }

	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
		for j := range vis[i] {
			vis[i][j] = -1
		}
	}

	q := []pair{}
	q = append(q, pair{0, 0, k})
	vis[0][0] = k
	step := 0

	if m == 1 && n == 1 {
		return step
	}

	for len(q) > 0 {
		step++
		tmp := q
		q = []pair{}
		for _, p := range tmp {
			for _, dir := range dirs {
				nx, ny := p.x+dir.x, p.y+dir.y
				ck := p.count
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if nx == m-1 && ny == n-1 {
						return step
					}
					if grid[nx][ny] != 0 {
						ck--
					}
					if ck >= 0 {
						if vis[nx][ny] == -1 || (vis[nx][ny] != -1 && ck > vis[nx][ny]) {
							q = append(q, pair{nx, ny, ck})
							vis[nx][ny] = ck
						}
					}
				}
			}
		}
	}
	return -1
}
