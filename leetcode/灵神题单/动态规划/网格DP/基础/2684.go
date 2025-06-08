package main

// dfs(i,j) = max(dfs(i-1,j),dfs(i,j),dfs(i+1,j)) + 1

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for _, row := range grid {
		row[0] *= -1 // 入队标记
	}
	for j := 0; j < n-1; j++ {
		ok := false
		for i := 0; i < m; i++ {
			if grid[i][j] > 0 { // 不在队列中
				continue
			}
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if grid[k][j+1] > -grid[i][j] {
					grid[k][j+1] *= -1 // 入队标记
					ok = true
				}
			}
		}
		if !ok { // 无法再往右走了
			return j
		}
	}
	return n - 1
}
