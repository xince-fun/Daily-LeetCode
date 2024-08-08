package main

import "sort"

// 分组

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i, x := range arr {
		m[x] = append(m[x], i)
	}
	return RangeFreqQuery{
		m: m,
	}
}

func (t *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := t.m[value]
	return sort.SearchInts(arr, left) - sort.SearchInts(arr, right+1)
}
