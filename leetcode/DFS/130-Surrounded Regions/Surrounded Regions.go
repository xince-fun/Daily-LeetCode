package main

/*
func solve(board [][]byte) {
	r, c := len(board), len(board[0]) // r -> row c -> column

	for i := 0; i < r; i++ {
		dfs(board, r, c, i, 0)
		dfs(board, r, c, i, c-1)
	}
	for i := 1; i < c-1; i++ {
		dfs(board, r, c, 0, i)
		dfs(board, r, c, r-1, i)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			switch board[i][j] {
			case 'A':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, r, c, x, y int) {
	if x < 0 || x >= r || y < 0 || y >= c || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'

	dfs(board, r, c, x-1, y)
	dfs(board, r, c, x+1, y)
	dfs(board, r, c, x, y-1)
	dfs(board, r, c, x, y+1)
}
*/

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'

		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 1; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'A':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}
