package main

import "fmt"

func orangesRotting(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	fresh := 0
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	type pair struct{ i, j int }
	q := make([]pair, 0)
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
			p := q[0]
			q = q[1:]
			for _, d := range direction {
				x, y := p.i+d[0], p.j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					q = append(q, pair{x, y})
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
	return
}

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
