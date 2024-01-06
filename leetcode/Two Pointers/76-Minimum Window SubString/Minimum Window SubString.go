package main

import (
	"math"
)

func minWindow(s string, t string) (ans string) {
	left := 0
	minLen := math.MaxInt
	m := map[rune]int{}
	for _, c := range t {
		m[c]++
	}
	for right, c := range s {
		m[c]--
		for check(m) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}
			m[rune(s[left])]++
			left++
		}
	}
	return ans
}

func check(m map[rune]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
