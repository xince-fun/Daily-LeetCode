package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	count := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			count[grid[i][j]]++
		}
	}

	var a, b int
	for i := 1; i <= n*n; i++ {
		if c, exists := count[i]; !exists {
			b = i
		} else if c == 2 {
			a = i
		}
	}
	return []int{a, b}
}
