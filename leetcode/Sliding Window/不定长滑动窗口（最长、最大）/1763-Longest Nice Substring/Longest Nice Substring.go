package main

import "strings"

func longestNiceSubstring1(s string) (ans string) {
	check := func(s string) bool {
		m := make(map[byte]int)
		for _, c := range s {
			m[byte(c)]++
		}
		for _, c := range s {
			if m[byte(c)] == 0 || m[byte(c^32)] == 0 {
				return false
			}
		}
		return true
	}

	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if j-i+1 > len(ans) && check(s[i:j+1]) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

func longestNiceSubstring2(s string) (ans string) {
	n := len(s)
	cnt := make([][128]int, n+1)
	for i := 0; i < n; i++ {
		c := s[i]
		cnt[i+1] = cnt[i]
		cnt[i+1][c-'A']++
	}

	check := func(a, b [128]int) bool {
		for i := 0; i < 26; i++ {
			low, up := b[i]-a[i], b[i+32]-a[i+32]
			if low != 0 && up == 0 {
				return false
			}
			if low == 0 && up != 0 {
				return false
			}
		}
		return true
	}

	idx, len := -1, 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if j-i+1 <= len {
				continue
			}
			if check(cnt[i], cnt[j+1]) {
				idx = i
				len = j - i + 1
			}
		}
	}
	if idx == -1 {
		return ans
	}
	return s[idx : idx+len]
}

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if (s[i] < 97 && !strings.Contains(s, string(s[i]+32))) || (s[i] >= 97 && !strings.Contains(s, string(s[i]-32))) {
			s1, s2 := longestNiceSubstring(s[:i]), longestNiceSubstring(s[i+1:])
			if len(s1) >= len(s2) {
				return s1
			}
			return s2
		}
	}
	return s
}
