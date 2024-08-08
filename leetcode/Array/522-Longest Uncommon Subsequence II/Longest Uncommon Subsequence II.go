package main

import "sort"

func findLUSlength(strs []string) int {
	n := len(strs)
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	for i := 0; i < n; i++ {
		found := false
		for j := 0; j < n; j++ {
			if j != i && isSubSeq(strs[i], strs[j]) {
				found = true
				break
			}
		}
		if found {
			return len(strs[i])
		}
	}
	return -1
}

func isSubSeq(a, b string) bool {
	i := 0
	for _, c := range b {
		if byte(c) == a[i] {
			i++
			if i == len(a) {
				return true
			}
		}
	}
	return false
}
