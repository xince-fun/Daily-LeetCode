package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(maze [][]byte, entrance []int) (ans int) {
	type pair struct{ x, y int }
	q := []pair{}
	m, n := len(maze), len(maze[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	q = append(q, pair{entrance[0], entrance[1]})
	vis[entrance[0]][entrance[1]] = true
	count := 0
	for len(q) > 0 {
		length := len(q)
		count++
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] && maze[nx][ny] == '.' {
					if (nx == 0 || nx == m-1 || ny == 0 || ny == n-1) && maze[nx][ny] == '.' {
						return count
					}
					q = append(q, pair{nx, ny})
					vis[nx][ny] = true
				}
			}
		}
	}
	return -1
}
