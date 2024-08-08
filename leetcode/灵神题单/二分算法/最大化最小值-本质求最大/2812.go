package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 最大化最小值
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	type pair struct{ x, y, t int }

	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = -1
		}
	}

	q := []pair{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		x, y, t := e.x, e.y, e.t
		if dis[x][y] != -1 {
			continue
		}
		dis[x][y] = t
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx >= 0 && nx < n && ny >= 0 && ny < n && dis[nx][ny] == -1 {
				q = append(q, pair{nx, ny, t + 1})
			}
		}
	}

	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = -1
		}
	}

	check := func(mid int) bool {
		type pair struct{ x, y int }
		q := []pair{}
		q = append(q, pair{0, 0})
		if dis[0][0] < mid {
			return false
		}
		for len(q) > 0 {
			e := q[0]
			q = q[1:]
			if e.x == n-1 && e.y == n-1 {
				return true
			}
			if vis[e.x][e.y] == mid {
				continue
			}
			vis[e.x][e.y] = mid
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < n && ny >= 0 && ny < n && vis[nx][ny] != mid && dis[nx][ny] >= mid {
					q = append(q, pair{nx, ny})
				}
			}
		}
		return false
	}

	left, right := -1, 2*n
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
