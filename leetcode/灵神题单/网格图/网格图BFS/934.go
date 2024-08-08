package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, y int }

	q := []pair{}

	for i := 0; i < m; i++ {
		found := false
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j})
				grid[i][j] = 2
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for _, dir := range dirs {
			nx, ny := e.x+dir.x, e.y+dir.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
				q = append(q, pair{nx, ny})
				grid[nx][ny] = 2
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j})
				grid[i][j] = 3
			}
		}
	}

	count := 0
	for len(q) > 0 {
		length := len(q)
		add := false
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if grid[nx][ny] == 2 {
						return count
					}
					if grid[nx][ny] == 0 {
						add = true
						grid[nx][ny] = 3
						q = append(q, pair{nx, ny})
					}
				}
			}
		}
		if add {
			count++
		}
	}
	return count
}
