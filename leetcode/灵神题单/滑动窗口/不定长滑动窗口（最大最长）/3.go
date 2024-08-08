package main

func lengthOfLongestSubstring(s string) (ans int) {
	m := map[rune]int{}
	left := 0
	for right, c := range s {
		m[c]++
		for m[c] > 1 {
			m[rune(s[left])]--
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
