package main

func numberOfSubstrings(s string) (ans int) {
	cnt := [3]int{}
	left := 0
	for _, c := range s {
		cnt[c-'a']++
		for cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return
}
