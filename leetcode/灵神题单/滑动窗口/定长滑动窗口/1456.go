package main

import "fmt"

func maxVowels1(s string, k int) (ans int) {
	count := 0
	left := 0
	for right, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
		if right-left >= k {
			if s[left] == 'a' || s[left] == 'e' || s[left] == 'i' || s[left] == 'o' || s[left] == 'u' {
				count--
			}
			left++
		}
		ans = max(ans, count)
	}
	return ans
}

func maxVowels(s string, k int) (ans int) {
	vowel := 0

	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			vowel++
		}
		if i < k-1 { // 窗口大小不足 k
			continue
		}
		ans = max(ans, vowel)
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			vowel--
		}
	}
	return
}

func main() {
	str := "weallloveyou"
	k := 7
	fmt.Println(maxVowels(str, k))
}
