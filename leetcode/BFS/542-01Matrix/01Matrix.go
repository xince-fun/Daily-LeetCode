package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func updateMatrix(mat [][]int) [][]int {
	type pair struct{ x, y int }
	q := []pair{}
	m, n := len(mat), len(mat[0])

	vis := make([][]bool, m)
	dis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
		dis[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, pair{i, j})
				vis[i][j] = true
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for _, dir := range dirs {
			nx, ny := e.x+dir.x, e.y+dir.y
			if nx >= 0 && nx <= m-1 && ny >= 0 && ny <= n-1 && !vis[nx][ny] {
				dis[nx][ny] = dis[e.x][e.y] + 1
				vis[nx][ny] = true
				q = append(q, pair{nx, ny})
			}
		}
	}
	return dis
}
