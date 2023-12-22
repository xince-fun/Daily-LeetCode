package main

func findPeakGrid(mat [][]int) (ans []int) {
	type pair struct{ i, j int }
	m, n := len(mat), len(mat[0])
	// 先看横着的
	for i := 1; i < m-1; i++ {
		left, right := -1, n
		for left+1 < right {
			mid := left + (right-left)/2
			if mat[i][mid] > mat[i][mid+1] {
				right = mid
			} else {
				left = mid
			}
		}
		if mat[i][right] > mat[i-1][right] && mat[i+1][right] > mat[i][right] {
			return []int{i, right}
		}
	}
	return []int{}
}
