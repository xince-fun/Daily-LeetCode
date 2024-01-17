package main

import "sort"

func beautifulIndices(s string, a string, b string, k int) (ans []int) {
	posA := kmp(s, a)
	posB := kmp(s, b)

	for _, i := range posA {
		bi := sort.SearchInts(posB, i)
		// 离 i 最近的 j 是 posB[bi] 或 posB[bi-1]
		if bi < len(posB) && posB[bi]-i <= k || bi > 0 && i-posB[bi-1] <= k {
			ans = append(ans, i)
		}
	}
	return
}

func kmp(text, pattern string) (pos []int) {
	m := len(pattern)
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for i, v := range text {
		for cnt > 0 && pattern[cnt] != byte(v) {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == byte(v) {
			cnt++
		}
		if cnt == m {
			pos = append(pos, i-m+1)
			cnt = pi[cnt-1]
		}
	}
	return
}

// abs returns the absolute difference between two integers
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// main function for testing
func main() {
	s := "isawsquirrelnearmysquirrelhouseohmy"
	a := "my"
	b := "squirrel"
	k := 15

	indices := beautifulIndices(s, a, b, k)
	// Use indices as needed, e.g., print them
	for _, index := range indices {
		println(index)
	}

	// You can also test the beautifulIndicesOld function similarly
}
