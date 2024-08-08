package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting1(grid [][]int) (ans int) {
	type pair struct{ x, y int }

	m, n := len(grid), len(grid[0])
	q := []pair{}

	fresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, pair{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	for len(q) > 0 {
		length := len(q)
		rotten := false
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]

			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					q = append(q, pair{nx, ny})
					grid[nx][ny] = 2
					fresh--
					rotten = true
				}
			}
		}
		if rotten {
			ans++
		}
	}

	if fresh > 0 {
		return -1
	}
	return ans
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	type pair struct{ x, y int }

	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, pair{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	ans := -1
	for len(q) > 0 {
		ans++
		tmp := q
		q = []pair{}
		for _, p := range tmp { // 已经腐烂的
			for _, d := range dirs {
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
					fresh--
					grid[i][j] = 2
					q = append(q, pair{i, j})
				}
			}
		}
	}
	if fresh > 0 {
		return -1
	}
	return max(ans, 0)
}
