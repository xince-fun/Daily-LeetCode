package main

import (
	"fmt"
	"sort"
)

func swimInWater(grid [][]int) int {
	n := len(grid)
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	// 相当于最小化路径最大值
	return sort.Search(n*n, func(t int) bool {
		if t < grid[0][0] {
			return false
		}
		g := make([][]bool, n)
		for i := 0; i < n; i++ {
			g[i] = make([]bool, n)
			for j := 0; j < n; j++ {
				g[i][j] = t >= grid[i][j]
			}
		}
		// 从(0,0)到(n-1,n-1)
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		visited[0][0] = true
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{0, 0})

		for len(queue) > 0 {
			point := queue[0]
			queue = queue[1:]

			if point[0] == n-1 && point[1] == n-1 {
				return true
			}
			for _, dir := range directions {
				nRow, nCol := point[0]+dir[0], point[1]+dir[1]
				if nRow >= 0 && nRow < n && nCol >= 0 && nCol < n && !visited[nRow][nCol] && g[nRow][nCol] {
					queue = append(queue, [2]int{nRow, nCol})
					visited[nRow][nCol] = true
				}
			}
		}
		return false
	})
}

func main() {
	grid := [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}
	fmt.Println(swimInWater(grid))
}
