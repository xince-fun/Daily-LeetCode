package main

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 1; j < n-1; j++ {
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
