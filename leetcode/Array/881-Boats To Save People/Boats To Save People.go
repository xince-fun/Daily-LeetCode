package main

import (
	"slices"
)

func numRescueBoats(people []int, limit int) (ans int) {
	slices.SortFunc(people, func(a, b int) int {
		return a - b
	})
	left, right := 0, len(people)-1
	for left <= right {
		lNum := people[left]
		rNum := people[right]
		if lNum+rNum > limit {
			right--
		} else {
			left++
			right--
		}
		ans++
	}
	return
}
