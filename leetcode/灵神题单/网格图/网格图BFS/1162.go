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

	ans := -1
	for len(q) > 0 {
		ans++
		tmp := q
		q = []pair{}
		for _, p := range tmp {
			for _, dir := range dirs {
				nx, ny := p.x+dir.x, p.y+dir.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == 0 {
					q = append(q, pair{nx, ny})
					grid[nx][ny] = 2
				}
			}
		}
	}
	return ans
}
