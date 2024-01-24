package main

import "fmt"

func longestSubstring1(s string, k int) (ans int) {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] < k {
			l := longestSubstring1(s[0:i], k)
			r := longestSubstring1(s[i+1:], k)
			return max(l, r)
		}
	}
	return len(s)
}

func longestSubstring(s string, k int) (ans int) {
	// 固定窗口内的字母数
	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	// 所有的字母数 循环开始
	n := len(hash)
	for i := 1; i <= n; i++ {
		// 现在固定窗口内 只能出现的字母数为i
		left := 0
		cnt := 0
		m := make(map[byte]int)
		for right, x := range s {
			m[byte(x)]++
			if m[byte(x)] == k {
				cnt++
			}
			for len(m) > i {
				if m[s[left]] == k {
					cnt--
				}
				m[s[left]]--
				if m[s[left]] == 0 {
					delete(m, s[left])
				}
				left++
			}
			if cnt == i {
				ans = max(ans, right-left+1)
			}
		}
	}
	return
}

func longestSubstring2(s string, k int) (ans int) {
	// 固定窗口内的字母数
	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	// 所有的字母数 循环开始
	n := len(hash)
	for i := 1; i <= n; i++ {
		// 现在固定窗口内 只能出现的字母数为i
		left := 0
		cnt := 0
		cnt1 := 0
		// m := make(map[byte]int)
		m := [26]int{}
		for right, x := range s {
			if m[x-'a'] == 0 {
				cnt1++
			}
			m[x-'a']++
			if m[x-'a'] == k {
				cnt++
			}
			for cnt1 > i {
				if m[s[left]-'a'] == k {
					cnt--
				}
				m[s[left]-'a']--
				if m[s[left]-'a'] == 0 {
					cnt1--
				}
				left++
			}
			if cnt == i {
				ans = max(ans, right-left+1)
			}
		}
	}
	return
}

func main() {
	s := "ababbc"
	k := 2
	fmt.Println(s, k)
	fmt.Println(longestSubstring(s, k))
}
