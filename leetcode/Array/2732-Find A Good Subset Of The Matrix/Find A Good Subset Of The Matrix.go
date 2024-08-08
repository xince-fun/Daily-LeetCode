package main

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	m := len(grid)
	bitmasks := make([]int, m)
	for i := 0; i < m; i++ {
		bitmask := 0
		for _, val := range grid[i] {
			bitmask = (bitmask << 1) | val
		}
		bitmasks[i] = bitmask
	}
	for i := 0; i < m; i++ {
		// size: 1
		if bitmasks[i] == 0 {
			return []int{i}
		}
		// size: 2
		for j := i + 1; j < m; j++ {
			if bitmasks[i]&bitmasks[j] == 0 {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
