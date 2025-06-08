package main

// 递归边界：当 i<0 或 j<0 时，返回 0，因为出界是没有价值的。也就是 dfs(−1,j)=0,dfs(i,−1)=0。

func jewelleryValue1(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p
		}
		*p = max(dfs(i, j-1), dfs(i-1, j)) + frame[i][j]
		return *p
	}
	return dfs(m-1, n-1)
}

func jewelleryValue2(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = frame[0][0]
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + frame[i][0]
	}
	for j := 1; j < n; j++ {
		f[0][j] = f[0][j-1] + frame[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1]) + frame[i][j]
		}
	}
	return f[m-1][n-1]
}

// 当 i=0 或 j=0 时，会出现负数下表，这种定义方式没有状态能表示递归边界
// 在 f 数组的上边和左边加一排， 把原来的 f[i+1] 改成 f[i+1], f[i-1] 改成 f[i]
// f[i][0] f[0][j] 表示出界的状态

// 初始值 f[i][0] = 0, f[0][j] = 0 ( 翻译自 dfs(i, -1) = 0 和 dfs(-1, j) = 0 )

func jewelleryValue3(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, row := range frame {
		for j, x := range row {
			f[i+1][j+1] = max(f[i][j+1], f[i+1][j]) + x
		}
	}
	return f[m][n]
}

func jewelleryValue4(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	f := make([]int, n)
	f[0] = frame[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				f[j] = f[j-1] + frame[i][j]
			} else if j == 0 && i > 0 {
				f[j] = f[j] + frame[i][j]
			} else if j > 0 {
				f[j] = max(f[j], f[j-1]) + frame[i][j]
			}
		}
	}
	return f[n-1]
}

func jewelleryValue5(frame [][]int) int {
	m, n := len(frame), len(frame[0])
	f := [2][]int{make([]int, n+1), make([]int, n+1)}
	for i, row := range frame {
		for j, x := range row {
			f[(i+1)%2][j+1] = max(f[i%2][j+1], f[(i+1)%2][j]) + x
		}
	}
	return f[m%2][n]
}

func jewelleryValue6(grid [][]int) int {
	n := len(grid[0])
	f := make([]int, n+1)
	for _, row := range grid {
		for j, x := range row {
			f[j+1] = max(f[j], f[j+1]) + x
		}
	}
	return f[n]
}

// 原地修改
// 直接用grid[0] 作为 f 数组
// f[i][j] = max(f[i][j-1], f[i-1][j]) + grid[i][j]
// i=0 和 j=0 的情况单独计算

func jewelleryValue(grid [][]int) int {
	n := len(grid[0])
	f := grid[0]
	for j := 1; j < n; j++ {
		f[j] += f[j-1]
	}
	for _, row := range grid[1:] {
		f[0] += row[0]
		for j := 1; j < n; j++ {
			f[j] = max(f[j], f[j-1]) + row[j]
		}
	}
	return f[n-1]
}
