package main

func maxVowels(s string, k int) (ans int) {
	m := map[byte]bool{}
	for _, c := range "aeiou" {
		m[byte(c)] = true
	}
	left, cnt := 0, 0
	for right, x := range s {
		if m[byte(x)] {
			cnt++
		}
		for right-left+1 > k {
			if m[s[left]] {
				cnt--
			}
			left++
		}
		ans = max(ans, cnt)
	}
	return
}
