package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	borders := []point{}

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	start := grid[row][col]
	var dfs func(i, j int)
	dfs = func(i, j int) {
		vis[i][j] = true
		isBorder := false

		for _, dir := range dirs {
			nx, ny := i+dir.x, j+dir.y
			if !(nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == start) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}

		}

		if isBorder {
			borders = append(borders, point{i, j})
		}
	}

	dfs(row, col)

	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}
