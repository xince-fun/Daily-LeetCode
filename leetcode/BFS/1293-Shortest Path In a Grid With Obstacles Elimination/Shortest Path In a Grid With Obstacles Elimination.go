package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, y, count int }

	vis := make([][]int, m)
	for i := range vis {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
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
		length := len(q)
		step++
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]

			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				currentK := e.count

				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if nx == m-1 && ny == n-1 {
						return step
					}

					if grid[nx][ny] != 0 {
						currentK--
					}
					if currentK >= 0 {
						if vis[nx][ny] == -1 || (vis[nx][ny] != -1 && currentK > vis[nx][ny]) {
							q = append(q, pair{nx, ny, currentK})
							vis[nx][ny] = currentK
						}
					}
				}
			}
		}
	}
	return -1
}
