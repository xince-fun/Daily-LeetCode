package main

func longestSemiRepetitiveSubstring1(s string) (ans int) {
	last := s[0]
	left := 0
	count := 0
	for right, c := range s[1:] {
		if c == rune(last) {
			count++
		}
		last = byte(c)
		for count > 1 {
			leftC := s[left]
			leftC1 := s[left+1]
			if leftC == leftC1 {
				count--
			}
			left++
		}
		ans = max(ans, right-left+2)
	}
	return
}

func longestSemiRepetitiveSubstring(s string) int {
	ans, left, same := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			same++
			if same > 1 {
				left++
				for s[left] != s[left-1] {
					left++
				}
				same = 1
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
