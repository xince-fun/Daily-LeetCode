package main

import "slices"

func onesCount(x int) (ones int) {
	for ; x > 0; x &= (x - 1) {
		ones++
	}
	return
}

func sortByBits1(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		lena, lenb := onesCount(a), onesCount(b)
		if lena != lenb {
			return lena - lenb
		}
		return a - b
	})
	return arr
}

var bits = [1e4 + 1]int{}

func init() {
	for i := 1; i <= 1e4; i++ {
		bits[i] = bits[i>>1] + i&1
	}
}

func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		cx, cy := bits[a], bits[b]
		if cx != cy {
			return cx - cy
		}
		return a - b
	})
	return arr
}
