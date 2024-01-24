package main

import (
	"fmt"
)

func numberOfSubstrings(s string) (ans int) {
	cnt := [3]int{}
	left := 0
	for _, x := range s {
		cnt[x-'a']++
		for cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return
}

func main() {
	s := "aaacb"
	fmt.Println(numberOfSubstrings(s))
}
