package main

import (
	"fmt"
	"sort"
)

func maximumRemovals(s string, p string, removable []int) int {
	// 自己写 [left, right)
	check := func(k int) bool {
		del := make([]bool, len(s))
		for _, i := range removable[:k] {
			del[i] = true
		}
		i, j := 0, 0
		for i < len(s) && j < len(p) {
			if !del[i] && s[i] == p[j] {
				j++
			}
			i++
		}
		return j == len(p)
	}

	left := 0
	right := len(removable) + 1
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// 循环不变量
// left-1 回答一定是「是」
// right 的回答一定是「否」

func maximumRemovals1(s string, p string, removable []int) int {
	return sort.Search(len(removable), func(k int) bool {
		del := make([]bool, len(s))
		for _, i := range removable[:k+1] {
			del[i] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				if j++; j == len(p) {
					return false
				}
			}
		}
		return true
	})
}

func main() {
	s := "abcbddddd"
	p := "abcd"
	removable := []int{3, 2, 1, 4, 5, 6}
	fmt.Println(maximumRemovals(s, p, removable))
}
