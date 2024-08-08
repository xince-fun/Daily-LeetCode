package main

func findPeakGrid(mat [][]int) []int {
	left, right := -1, len(mat)-1
	for left+1 < right {
		i := left + (right-left)/2
		j := indexOfMax(mat[i])
		if mat[i][j] > mat[i+1][j] {
			right = i // 峰顶行号<=i
		} else {
			left = i // 峰顶行号>=i
		}
	}
	return []int{right, indexOfMax(mat[right])}
}

func indexOfMax(a []int) (idx int) {
	for i, x := range a {
		if x > a[idx] {
			idx = i
		}
	}
	return
}
