package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rowArr, colArr := make([]int, m), make([]int, n)
	type pair struct{ x, y int }
	loc := make(map[int]pair)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			loc[mat[i][j]] = pair{i, j}
		}
	}
	for i, a := range arr {
		p := loc[a]
		rowArr[p.x]++
		colArr[p.y]++
		if rowArr[p.x] == n || colArr[p.y] == m {
			return i
		}
	}
	return 0
}
