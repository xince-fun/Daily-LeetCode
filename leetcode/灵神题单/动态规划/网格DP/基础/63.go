package main

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[1][0] = 1
	for i, row := range obstacleGrid {
		for j, x := range row {
			if x == 0 {
				f[i+1][j+1] = f[i][j+1] + f[i+1][j]
			} else {
				f[i+1][j+1] = 0
			}
		}
	}
	return f[m][n]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := [2][]int{make([]int, n+1), make([]int, n+1)}
	f[0][1] = 1
	for i, row := range obstacleGrid {
		for j, x := range row {
			if x == 0 {
				f[(i+1)%2][j+1] = f[i%2][j+1] + f[(i+1)%2][j]
			} else {
				f[(i+1)%2][j+1] = 0
			}
		}
	}
	return f[m%2][n]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	f := make([]int, n+1)
	f[1] = 1
	for _, row := range obstacleGrid {
		for j, x := range row {
			if x == 0 {
				f[j+1] = f[j+1] + f[j]
			} else {
				f[j+1] = 0
			}
		}
	}
	return f[n]
}
