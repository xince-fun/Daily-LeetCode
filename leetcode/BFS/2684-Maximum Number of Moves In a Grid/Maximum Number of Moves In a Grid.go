package main

var dirs = []struct{ x, y int }{{0, 1}, {1, 1}, {-1, 1}}

func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	type pair struct{ x, y int }
	q := []pair{}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		q = append(q, pair{i, 0})
		vis[i][0] = true
	}

	for len(q) > 0 {
		length := len(q)
		moved := false
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] && grid[nx][ny] > grid[e.x][e.y] {
					q = append(q, pair{nx, ny})
					vis[nx][ny] = true
					moved = true
				}
			}
		}
		if moved {
			ans++
		}
	}
	return
}
