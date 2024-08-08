package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (res [][]int) {
	m, n := len(heights), len(heights[0])

	atlantic := make([][]bool, m)
	pacific := make([][]bool, m)
	for i := 0; i < m; i++ {
		atlantic[i] = make([]bool, n)
		pacific[i] = make([]bool, n)
	}

	var dfs func(x, y int, ocean [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}

		ocean[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}

	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return
}
