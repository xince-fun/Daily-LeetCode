package main

// E23
// 2X2
// 12S

// 考虑从左上角出发，到达右下角的,只能往右，往下，往右下
// dfs(i,j) = max(dfs(i-1,j), dfs(i,j-1), dfs(i-1,j-1)) + grid[i][j] 表示最大的值

// f[][][2] 0表示最大值， 1表示方案数
func pathsWithMaxScore(board []string) []int {
	const mod = 1_000_000_007
	n := len(board)
	f := make([][][]int, n+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, 2)
		}
		f[i][0][0] = -1
	}
	for j := 1; j <= n; j++ {
		f[0][j][0] = -1
	}
	f[0][0][0], f[0][0][1] = 0, 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				f[i+1][j+1][0], f[i+1][j+1][1] = -1, 0
			} else {
				cnt, maxn := 0, -1
				for _, x := range [][]int{f[i][j+1], f[i+1][j], f[i][j]} {
					if x[0] > maxn {
						maxn = x[0]
						cnt = x[1]
					} else if x[0] == maxn {
						cnt += x[1]
					}
				}
				num := 0
				if board[i][j] != 'E' && board[i][j] != 'S' {
					num = int(board[i][j] - '0')
				}
				if maxn == -1 {
					f[i+1][j+1][0], f[i+1][j+1][1] = -1, cnt%mod
				} else {
					f[i+1][j+1][0], f[i+1][j+1][1] = maxn+num, cnt%mod
				}

			}
		}
	}
	if f[n][n][0] == -1 {
		return []int{0, 0}
	}
	return []int{f[n][n][0] % mod, f[n][n][1] % mod}
}
