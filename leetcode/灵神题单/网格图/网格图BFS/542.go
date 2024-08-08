package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	vis := make([][]bool, m)
	dis := make([][]int, m)
	for i := range dis {
		vis[i] = make([]bool, n)
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}

	type pair struct{ x, y, t int }
	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, pair{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		if dis[e.x][e.y] != -1 {
			continue
		}
		dis[e.x][e.y] = e.t
		for _, dir := range dirs {
			nx, ny := e.x+dir.x, e.y+dir.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && dis[nx][ny] == -1 {
				q = append(q, pair{nx, ny, e.t + 1})
			}
		}
	}
	return dis
}
