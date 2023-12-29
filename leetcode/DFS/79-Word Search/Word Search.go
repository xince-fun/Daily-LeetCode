package main

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if i >= m || i < 0 || j >= n || j < 0 || board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = ' '

		res := dfs(i+1, j, index+1) || dfs(i-1, j, index+1) || dfs(i, j-1, index+1) || dfs(i, j+1, index+1)
		board[i][j] = temp
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
