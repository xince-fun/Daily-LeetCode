package main

var dirs = []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(x, y int) {
		cnt := 0

		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if board[nx][ny] == 'M' {
				cnt++
			}
		}
		if cnt > 0 {
			board[x][y] = byte(cnt + '0')
		} else {
			board[x][y] = 'B'
			for _, dir := range dirs {
				nx, ny := x+dir.x, y+dir.y
				if nx < 0 || nx >= m || ny < 0 || ny >= n || board[nx][ny] != 'E' {
					continue
				}
				dfs(nx, ny)
			}
		}
	}

	x, y := click[0], click[1]

	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(x, y)
	}
	return board

}
