package main

func findAnagrams(s string, p string) (ans []int) {
	if len(p) > len(s) {
		return
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range p {
		cnt1[p[i]-'a']++
		cnt2[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if cnt1 == cnt2 {
			ans = append(ans, i)
		}
		cnt2[s[i+len(p)]-'a']++
		cnt2[s[i]-'a']--
	}
	if cnt1 == cnt2 {
		ans = append(ans, len(s)-len(p))
	}
	return
}
