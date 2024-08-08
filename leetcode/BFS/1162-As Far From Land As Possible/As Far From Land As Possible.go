package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxDistance(grid [][]int) int {
	type pair struct{ x, y int }
	m, n := len(grid), len(grid[0])

	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j})
				grid[i][j] = 2
			}
		}
	}

	if len(q) == 0 || len(q) == m*n {
		return -1
	}

	count := 0
	for len(q) > 0 {
		length := len(q)
		isAdd := false
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
					q = append(q, pair{nx, ny})
					grid[nx][ny] = 2
					isAdd = true
				}
			}
		}
		if isAdd {
			count++
		}
	}
	return count
}
