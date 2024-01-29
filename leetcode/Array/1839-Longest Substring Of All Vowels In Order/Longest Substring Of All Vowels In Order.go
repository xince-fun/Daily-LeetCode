package main

func longestBeautifulSubstring1(word string) (ans int) {
	check := func(a, b byte) bool {
		if a == b {
			return true
		}
		if a == 'a' {
			return b == 'e'
		} else if a == 'e' {
			return b == 'i'
		} else if a == 'i' {
			return b == 'o'
		} else if a == 'o' {
			return b == 'u'
		}
		return false
	}
	for i, n := 0, len(word); i < n; {
		if word[i] != 'a' {
			i++
			continue
		}
		i0 := i
		for i++; i < n && check(word[i-1], word[i]); i++ {
		}
		if word[i-1] == 'u' {
			ans = max(ans, i-i0)
		}
	}
	return
}

func longestBeautifulSubstring(s string) (ans int) {
	const vowel = "aeiou"
	cur, sum := 0, 0
	for i, n := 0, len(s); i < n; {
		start := i
		ch := s[start]
		for i < n && s[i] == ch {
			i++
		}

		if ch != vowel[cur] {
			cur, sum = 0, 0
			if ch != vowel[0] {
				continue
			}
		}

		sum += i - start
		cur++
		if cur == 5 {
			if sum > ans {
				ans = sum
			}
			cur, sum = 0, 0
		}
	}
	return
}
