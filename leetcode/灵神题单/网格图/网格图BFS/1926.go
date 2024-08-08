package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit1(maze [][]byte, entrance []int) (ans int) {
	m, n := len(maze), len(maze[0])
	type pair struct{ x, y int }
	q := []pair{}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[entrance[0]][entrance[1]] = true
	q = append(q, pair{entrance[0], entrance[1]})
	for len(q) > 0 {
		ans++
		length := len(q)
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && !vis[nx][ny] && maze[nx][ny] == '.' {
					if (nx == 0 || nx == m-1 || ny == 0 || ny == n-1) && maze[nx][ny] == '.' {
						return ans
					}
					vis[nx][ny] = true
					q = append(q, pair{nx, ny})
				}
			}
		}
	}
	return -1
}

func nearestExit(maze [][]byte, entrance []int) (ans int) {
	m, n := len(maze), len(maze[0])
	type pair struct{ x, y int }
	q := []pair{}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[entrance[0]][entrance[1]] = true
	q = append(q, pair{entrance[0], entrance[1]})

	for len(q) > 0 {
		ans++
		tmp := q
		q = []pair{}
		for _, p := range tmp {
			for _, dir := range dirs {
				nx, ny := p.x+dir.x, p.y+dir.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && !vis[nx][ny] && maze[nx][ny] == '.' {
					if (nx == 0 || nx == m-1 || ny == 0 || ny == n-1) && maze[nx][ny] == '.' {
						return ans
					}
					vis[nx][ny] = true
					q = append(q, pair{nx, ny})
				}
			}
		}
	}
	return -1
}
