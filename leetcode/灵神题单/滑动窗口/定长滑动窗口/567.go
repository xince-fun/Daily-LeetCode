package main

func checkInclusion1(s1 string, s2 string) bool {
	m1 := map[rune]int{}
	for _, c := range s1 {
		m1[c]++
	}
	k := len(s1)
	m2 := map[rune]int{}
	for i, in := range s2 {
		m2[in]++
		if i >= k-1 {
			ok := true
			for c, v := range m2 {
				if m1[c] != v {
					ok = false
					break
				}
			}
			if ok {
				return true
			}
			out := rune(s2[i-k+1])
			m2[out]--
			if m2[out] == 0 {
				delete(m2, out)
			}
		}
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	cnt := [26]int{}
	for _, c := range s1 {
		cnt[c-'a']++
	}

	isValid := func(cnt [26]int) bool {
		for _, v := range cnt {
			if v != 0 {
				return false
			}
		}
		return true
	}

	m := len(s1)
	for i, in := range s2 {
		cnt[in-'a']--
		if i >= m-1 {
			if isValid(cnt) {
				return true
			}
			out := s2[i-m+1]
			cnt[out-'a']++
		}
	}
	return false
}
