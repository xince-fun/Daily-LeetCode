package main

import (
	"fmt"
	"sort"
)

func shipWithinDays(weights []int, days int) int {
	// [)
	return sort.Search(1e7, func(x int) bool {
		cnt, cur := 1, 0
		for _, weight := range weights {
			if x < weight {
				return false
			}
			if cur+weight > x {
				cnt++
				cur = 0
			}
			cur += weight
		}
		return cnt <= days
	})
}

func main() {
	weights := []int{3, 2, 2, 4, 1, 4}
	days := 3
	fmt.Println(shipWithinDays(weights, days))
}
