package main

func countBattleships1(board [][]byte) (ans int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				dfs(board, i, j)
				ans++
			}
		}
	}
	return
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '.' || board[i][j] == 'O' {
		return
	}

	board[i][j] = 'O'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func countBattleships(board [][]byte) (ans int) {
	for i, row := range board {
		for j, c := range row {
			if c == 'X' &&
				(j == 0 || row[j-1] != 'X') &&
				(i == 0 || board[i-1][j] != 'X') {
				ans++
			}
		}
	}
	return
}
