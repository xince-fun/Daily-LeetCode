package main

func findAnagrams1(s string, p string) []int {
	ans := make([]int, 0)
	m1 := map[rune]int{}
	for _, c := range p {
		m1[c]++
	}
	m2 := map[rune]int{}
	k := len(p)
	for i, in := range s {
		m2[in]++
		if i >= k-1 {
			ok := true
			for c, v := range m1 {
				if m2[c] != v {
					ok = false
				}
			}
			if ok {
				ans = append(ans, i-k+1)
			}
			out := rune(s[i-k+1])
			m2[out]--
			if m2[out] == 0 {
				delete(m2, out)
			}
		}
	}
	return ans
}

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	cnt := [26]int{}
	for _, c := range p {
		cnt[c-'a']++
	}

	isValid := func(cnt [26]int) bool {
		for _, num := range cnt {
			if num != 0 {
				return false
			}
		}
		return true
	}

	m := len(p)
	for i, in := range s {
		cnt[in-'a']--
		if i >= m-1 {
			if isValid(cnt) {
				ans = append(ans, i-m+1)
			}
			out := s[i-m+1]
			cnt[out-'a']++
		}
	}
	return ans
}
