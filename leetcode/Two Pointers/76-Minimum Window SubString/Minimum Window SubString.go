package main

import "math"

func minWindow(s string, t string) (ans string) {
	left := 0
	minLen := math.MaxInt
	m := map[rune]int{}
	for _, c := range t {
		m[c]++
	}
	for right, c := range s {
		if _, exists := m[c]; exists {
			m[c]--
		}
		isFound := true
		for _, v := range m {
			if v != 0 {
				isFound = false
				break
			}
		}
		if isFound {
			len := right - left + 1
			if len < minLen {
				minLen = len
				temp := append([]byte(nil), s[left:right+1]...)
				ans = string(temp)
			}
			// 开始缩小
			for left <= right {
				char := s[left]
				if _, exists := m[rune(char)]; exists {
					m[rune(char)]++
				}
				left++
			}
		}
	}
	return ans
}
