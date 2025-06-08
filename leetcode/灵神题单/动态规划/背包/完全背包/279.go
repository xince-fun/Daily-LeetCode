package main

import "math"

// f[i][j] = min(f[i-1][j],f[i][j-x]+1)

func numSquares1(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	m := len(nums)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt / 2
		}
	}
	f[0][0] = 0
	for i, x := range nums {
		for c := 0; c <= n; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-x]+1)
			}
		}
	}
	return f[m][n]
}

func numSquares(n int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
	}

	for _, x := range nums {
		for c := 0; c <= n; c++ {
			if c >= x {
				f[c] = min(f[c], f[c-x]+1)
			}
		}
	}
	return f[n]
}
