package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := m; i < n; i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']++
		cnt2[s2[i-m]-'a']--
	}
	return cnt1 == cnt2
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
