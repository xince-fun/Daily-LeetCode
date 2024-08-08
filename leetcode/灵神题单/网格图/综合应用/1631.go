package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	// bfs 求当前限定下能否到达
	type pair struct{ x, y int }
	check := func(mid int) bool {
		q := []pair{}
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		q = append(q, pair{0, 0})
		vis[0][0] = true
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == m-1 && p.y == n-1 {
				return true
			}
			for _, dir := range dirs {
				nx, ny := p.x+dir.x, p.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] &&
					abs(heights[nx][ny]-heights[p.x][p.y]) <= mid {
					vis[nx][ny] = true
					q = append(q, pair{nx, ny})
				}
			}
		}
		return false
	}

	left, right := -1, 1_000_000
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
