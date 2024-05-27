package main

import "sort"

func canEat1(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount)+1)
	for i, x := range candiesCount {
		sum[i+1] = sum[i] + x
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		maxC := (q[1] + 1) * q[2]
		minC := q[1]
		num := sum[q[0]+1] - sum[0]
		index := sort.SearchInts(sum, maxC)
		if index-1 >= q[0] && num > minC {
			ans[i] = true
		}
	}
	return ans
}

func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount)+1)
	for i, x := range candiesCount {
		sum[i+1] = sum[i] + x
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		minC := q[1] + 1
		maxC := minC * q[2]

		y1 := sum[q[0]] + 1
		y2 := sum[q[0]+1]
		ans[i] = !(minC > y2 || maxC < y1)
	}
	return ans
}
