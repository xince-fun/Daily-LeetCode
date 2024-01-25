package main

import (
	"fmt"
	"sort"
)

func minEatingSpeed1(piles []int, h int) int {
	maxPile := piles[0]
	for i := 0; i < len(piles); i++ {
		maxPile = max(maxPile, piles[i])
	}
	// [)
	return 1 + sort.Search(maxPile, func(k int) bool {
		k++
		cnt := 0
		for _, pile := range piles {
			cnt += (pile + k - 1) / k
		}
		return cnt <= h
	})
}

func minEatingSpeed(piles []int, h int) int {
	maxPile := piles[0]
	for i := 0; i < len(piles); i++ {
		maxPile = max(maxPile, piles[i])
	}
	// []
	cal := func(k int) bool {
		cnt := 0
		for _, pile := range piles {
			cnt += (pile + k - 1) / k
		}
		return cnt <= h
	}
	left := 1
	for left <= maxPile {
		// 循环不变量：
		// left-1一定为「否」
		// right+1一定为「是」
		mid := (left + maxPile) / 2
		if !cal(mid) {
			left = mid + 1
		} else {
			maxPile = mid - 1
		}
	}
	return maxPile + 1
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(piles, h))
}
