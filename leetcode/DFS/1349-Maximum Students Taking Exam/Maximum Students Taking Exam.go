package main

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	onPath := make([][]bool, m)
	for i := range onPath {
		onPath[i] = make([]bool, n)
	}

	valid := func(r, c int) bool {
		if seats[r][c] == '#' {
			return false
		}
		// 先看左边
		if c > 0 && (seats[r][c-1] != '#' && onPath[r][c-1] == true) {
			return false
		}
		// 右边
		if c < n-1 && (seats[r][c+1] != '#' && onPath[r][c+1] == true) {
			return false
		}
		// 左上角
		if r > 0 && c > 0 && (seats[r-1][c-1] != '#' && onPath[r-1][c-1] == true) {
			return false
		}
		// 右上角
		if r > 0 && c < 0 && (seats[r-1][c+1] != '#' && onPath[r-1][c+1] == true) {
			return false
		}
		return true
	}

	count, ans := 0, -1
	var dfs func(int)
	dfs = func(r int) {
		if r == m {
			ans = max(ans, count)
			return
		}
		for j, on := range onPath[r] {
			if !on {
				if valid(r, j) {
					count++
					onPath[r][j] = true
					dfs(r + 1)
					count--
					onPath[r][j] = false
				}
			}
		}
	}
	dfs(0)
	return ans
}
