package main

import "sort"

/*

 */

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(1e5, func(i int) bool {
		cnt := 0
		for _, q := range quantities {
			cnt += (q + i) / (i + 1)
		}
		return cnt <= n
	}) + 1
}
