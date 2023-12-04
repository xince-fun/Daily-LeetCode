package main

func lengthOfLongestSubstring(s string) int {
	ans, n := 0, len(s)
	left, right := 0, 0
	m := [26]int{}
	for right < n {
		index := s[right] - 'a'
		m[index]++
		right++
		for m[index] > 1 {
			index2 := s[left] - 'a'
			m[index2]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
