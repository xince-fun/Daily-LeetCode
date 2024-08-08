package main

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	first := true
	// direction: 0: down 1: up 2: left 3: right
	var dfs func(x, y, dir int)
	dfs = func(x, y, dir int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if vis[x][y] {
			return
		}

		if first {
			first = false
		} else {
			switch dir {
			case 0:
				if !(grid[x][y] == 2 || grid[x][y] == 5 || grid[x][y] == 6) {
					return
				}
			case 1:
				if !(grid[x][y] == 2 || grid[x][y] == 3 || grid[x][y] == 4) {
					return
				}
			case 2:
				if !(grid[x][y] == 1 || grid[x][y] == 4 || grid[x][y] == 6) {
					return
				}
			case 3:
				if !(grid[x][y] == 1 || grid[x][y] == 3 || grid[x][y] == 5) {
					return
				}
			}
		}

		vis[x][y] = true
		if grid[x][y] == 2 || grid[x][y] == 3 || grid[x][y] == 4 {
			dfs(x+1, y, 0)
		}
		if grid[x][y] == 2 || grid[x][y] == 5 || grid[x][y] == 6 {
			dfs(x-1, y, 1)
		}
		if grid[x][y] == 1 || grid[x][y] == 4 || grid[x][y] == 6 {
			dfs(x, y+1, 3)
		}
		if grid[x][y] == 1 || grid[x][y] == 3 || grid[x][y] == 5 {
			dfs(x, y-1, 2)
		}
	}

	if grid[0][0] == 1 || grid[0][0] == 6 { // 只能往右
		dfs(0, 0, 3)
	} else if grid[0][0] == 2 || grid[0][0] == 3 { // 只能往下
		dfs(0, 0, 0)
	} else if grid[0][0] == 4 {
		dfs(0, 0, 3)
		if vis[m-1][n-1] {
			return true
		} else {
			vis = make([][]bool, m)
			for i := 0; i < m; i++ {
				vis[i] = make([]bool, n)
			}
			dfs(0, 0, 0)
		}
	} else {
		return false
	}

	return vis[m-1][n-1]
}
