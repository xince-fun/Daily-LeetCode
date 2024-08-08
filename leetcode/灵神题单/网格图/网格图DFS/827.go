package main

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func largestIsland(grid [][]int) (ans int) {
	n := len(grid)

	area := []int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		grid[i][j] = len(area) + 2 // 记录(i,j)属于哪个岛
		size := 1
		for _, dir := range dirs {
			x, y := i+dir.x, j+dir.y
			if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
				size += dfs(x, y)
			}
		}
		return size
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				area = append(area, dfs(i, j))
			}
		}
	}
	if len(area) == 0 {
		return 1
	}

	s := map[int]bool{}
	for i, row := range grid {
		for j, x := range row {
			if x != 0 {
				continue
			}
			clear(s)
			newArea := 1
			for _, dir := range dirs {
				x, y := i+dir.x, j+dir.y
				if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] != 0 && !s[grid[x][y]] {
					s[grid[x][y]] = true
					newArea += area[grid[x][y]-2]
				}
			}
			ans = max(ans, newArea)
		}
	}
	if ans == 0 {
		return n * n
	}
	return
}
