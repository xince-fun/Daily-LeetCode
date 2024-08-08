package main

var dirs = []struct{ x, y int }{{0, 1}, {1, 1}, {-1, 1}}

func maxMoves1(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, y int }

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	q := []pair{}
	for i := 0; i < m; i++ {
		q = append(q, pair{i, 0})
		vis[i][0] = true
	}

	for len(q) > 0 {
		length := len(q)
		moved := false
		for i := 0; i < length; i++ {
			e := q[0]
			q = q[1:]
			for _, dir := range dirs {
				nx, ny := e.x+dir.x, e.y+dir.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] && grid[nx][ny] > grid[e.x][e.y] {
					q = append(q, pair{nx, ny})
					vis[nx][ny] = true
					moved = true
				}
			}
		}
		if moved {
			ans++
		}
	}
	return
}

func maxMoves2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, y int }

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	q := []pair{}
	for i := 0; i < m; i++ {
		q = append(q, pair{i, 0})
		vis[i][0] = true
	}

	ans := -1
	for len(q) > 0 {
		ans++
		tmp := q
		q = []pair{}
		for _, p := range tmp {
			for _, d := range dirs {
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < m && 0 <= j && j < n && !vis[i][j] && grid[i][j] > grid[p.x][p.y] {
					vis[i][j] = true
					q = append(q, pair{i, j})
				}
			}
		}
	}
	return max(ans, 0)
}

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vis := make([]int, m)
	q := make([]int, m)

	for i := range q {
		q[i] = i
	}

	for j := 0; j < n-1; j++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if vis[k] != j+1 && grid[k][j+1] > grid[i][j] {
					vis[k] = j + 1 // 第 k 行目前最右访问到了 j
					q = append(q, k)
				}
			}
		}
		if q == nil {
			return j
		}
	}
	return n - 1
}
