package main

import (
	"slices"
	"sort"
)

func numSmallerByFrequency1(queries []string, words []string) []int {
	cnt := []int{}
	ans := make([]int, len(queries))
	getMin := func(s string) int {
		count := 0
		minC := 'z' + 1
		for _, c := range s {
			if c < minC {
				minC = c
				count = 0
			}
			if c == minC {
				count++
			}
		}
		return count
	}

	for _, word := range words {
		cnt = append(cnt, getMin(word))
	}
	slices.Sort(cnt)
	// 找到 > x 的第一个数
	for i, q := range queries {
		x := getMin(q)
		ans[i] = len(cnt) - sort.SearchInts(cnt, x+1)
	}
	return ans
}
