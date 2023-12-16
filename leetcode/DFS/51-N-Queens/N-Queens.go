package main

import (
	"strings"
)

// 用一个长为n的数组col记录皇后的位置
// 即第i行的皇后在第col[i]列
// 那么col必须是一个0到n-1的排列
// col=[1,3,0,2] [2,0,3,1]

// 用r表示行号，c表示列号 (本质是一次函数求截距)
// 右上方向的皇后 r + c 是一个定值
// 左上方向的皇后 r - c 是一个定值

func solveNQueens1(n int) (ans [][]string) {
	col := make([]int, n)
	// 剩余可以枚举的列号
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
			// 对于每个col， 前面有c个'.'，后面有n-1-c个'.'
			temp := []string{}
			for _, c := range col {
				temp = append(temp, strings.Repeat(".", c)+"Q"+strings.Repeat(".", n-1-c))
			}
			ans = append(ans, temp)
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
	// 2*n-1为对角线的长度
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			// 对于每个col， 前面有c个'.'，后面有n-1-c个'.'
			temp := []string{}
			for _, c := range col {
				temp = append(temp, strings.Repeat(".", c)+"Q"+strings.Repeat(".", n-1-c))
			}
			ans = append(ans, temp)
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
