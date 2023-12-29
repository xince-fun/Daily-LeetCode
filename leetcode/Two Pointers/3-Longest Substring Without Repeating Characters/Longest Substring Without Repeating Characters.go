package main

func lengthOfLongestSubstring1(s string) int {
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

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	m := make(map[byte]int)
	for right := range s {
		index := s[right] - 'a'
		m[index]++
		for m[index] > 1 {
			index2 := s[left] - 'a'
			m[index2]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
