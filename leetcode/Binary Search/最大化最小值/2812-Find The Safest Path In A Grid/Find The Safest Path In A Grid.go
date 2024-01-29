package main

import (
	"fmt"
	"sort"
)

// 最大化最小值
func maximumSafenessFactor(grid [][]int) int {
	// 对于安全系数 d
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = -1
		}
	}

	q := make([][3]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, [3]int{i, j, 0})
			}
		}
	}
	// 求最近的1的距离
	// 多源bfs求最近的1的距离
	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		x, y, t := point[0], point[1], point[2]
		if dis[x][y] != -1 {
			continue
		}
		dis[x][y] = t
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && dis[nx][ny] == -1 {
				q = append(q, [3]int{nx, ny, t + 1})
			}
		}
	}
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = -1
		}
	}

	// 找到最小不满足的，-1
	return sort.Search(2*n, func(d int) bool {
		q := make([][2]int, 0)
		q = append(q, [2]int{0, 0})
		if dis[0][0] < d {
			return true
		}
		for len(q) > 0 {
			point := q[0]
			q = q[1:]
			x, y := point[0], point[1]
			if x == n-1 && y == n-1 {
				return false
			}
			if vis[x][y] == d {
				continue
			}
			vis[x][y] = d
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < n && ny >= 0 && ny < n && vis[nx][ny] != d && dis[nx][ny] >= d {
					q = append(q, [2]int{nx, ny})
				}
			}
		}
		return true
	}) - 1
}

// 最大化最小值
func maximumSafenessFactor1(grid [][]int) int {
	// 对于安全系数 d
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = -1
		}
	}

	q := make([][3]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, [3]int{i, j, 0})
			}
		}
	}
	// 求最近的1的距离
	// 多源bfs求最近的1的距离
	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		x, y, t := point[0], point[1], point[2]
		if dis[x][y] != -1 {
			continue
		}
		dis[x][y] = t
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && dis[nx][ny] == -1 {
				q = append(q, [3]int{nx, ny, t + 1})
			}
		}
	}
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = -1
		}
	}
	check := func(d int) bool {
		q := make([][2]int, 0)
		q = append(q, [2]int{0, 0})
		if dis[0][0] < d {
			return false
		}
		for len(q) > 0 {
			point := q[0]
			q = q[1:]
			x, y := point[0], point[1]
			if x == n-1 && y == n-1 {
				return true
			}
			if vis[x][y] == d {
				continue
			}
			vis[x][y] = d
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < n && ny >= 0 && ny < n && vis[nx][ny] != d && dis[nx][ny] >= d {
					q = append(q, [2]int{nx, ny})
				}
			}
		}
		return false
	}

	left, right := 0, 2*n
	for left < right {
		mid := (left + right + 1) / 2
		// 找出最大的
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	grid := [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}
	fmt.Println(maximumSafenessFactor(grid))
}
