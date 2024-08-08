package main

import "strings"

func solveNQueens1(n int) (ans [][]string) {
	col := make([]int, n)

	onPath := make([]bool, n)

	valid := func(r, c int) bool {
		for R := 0; R < r; R++ {
			C := col[R]
			if r+c == R+C || r-c == R-C {
				return false
			}
		}
		return true
	}

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			tmp := []string{}
			for _, c := range col {
				tmp = append(tmp, strings.Repeat(".", c)+"Q"+strings.Repeat(".", n-1-c))
			}
			ans = append(ans, tmp)
			return
		}
		for j, on := range onPath {
			if !on {
				if valid(r, j) {
					col[r] = j
					onPath[j] = true
					dfs(r + 1)
					onPath[j] = false
				}
			}
		}
	}
	dfs(0)
	return
}

// 优化
// 既然可以用r+c和r-c判断
// 干脆把r+c记录到一个布尔数组diag1中
// 在放皇后前判断，如果diag1[r+c]为真，则无法放皇后

// 对于r-c同理，用布尔数组diag2记录
// 为了避免负数需要加上 n-1

func solveNQueens(n int) (ans [][]string) {
	col := make([]int, n)
	onPath := make([]bool, n)

	diag1 := make([]bool, 2*n-1)
	// 为了避免负数需要加上n-1
	diag2 := make([]bool, 2*n-1)

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			tmp := []string{}
			for _, c := range col {
				tmp = append(tmp, strings.Repeat(".", c)+"Q"+strings.Repeat(".", n-1-c))
			}
			ans = append(ans, tmp)
			return
		}
		for c, on := range onPath {
			rc := r - c + n - 1
			if !on && !diag1[r+c] && !diag2[rc] {
				col[r] = c
				onPath[c], diag1[r+c], diag2[rc] = true, true, true
				dfs(r + 1)
				onPath[c], diag1[r+c], diag2[rc] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
