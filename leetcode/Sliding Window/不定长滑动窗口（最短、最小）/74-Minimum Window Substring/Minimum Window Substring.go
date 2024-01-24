package main

func minWindow1(s string, t string) (ans string) {
	hash := make(map[byte]int)
	for _, c := range t {
		hash[byte(c)]++
	}
	check := func(m map[byte]int) bool {
		for _, v := range m {
			if v > 0 {
				return false
			}
		}
		return true
	}
	left := 0
	minLen := len(s) + 1
	for right, c := range s {
		hash[byte(c)]--
		for check(hash) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}
			hash[s[left]]++
			left++
		}
	}
	return
}

func minWindow2(s, t string) (ans string) {
	if len(t) > len(s) {
		return
	}

	dCnt := [128]int{}
	for i := 0; i < len(t); i++ {
		dCnt[t[i]]++
	}

	cCnt := [128]int{}
	differ := len(t)

	left := 0
	minLen := len(s) + 1

	for right := 0; right < len(s); right++ {
		if dCnt[s[right]] == 0 {
			continue
		}

		cCnt[s[right]]++
		if cCnt[s[right]] <= dCnt[s[right]] {
			differ--
		}

		for differ == 0 {
			if dCnt[s[left]] == 0 {
				left++
				continue
			}

			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}

			cCnt[s[left]]--
			if cCnt[s[left]] < dCnt[s[left]] {
				differ++
			}
			left++
		}
	}
	return
}

func minWindow(s, t string) (ans string) {
	if len(t) > len(s) {
		return
	}

	dCnt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		dCnt[t[i]]++
	}

	differ := len(dCnt)
	left := 0
	minLen := len(s) + 1

	for right := 0; right < len(s); right++ {
		if _, ok := dCnt[s[right]]; ok {
			dCnt[s[right]]--
			if dCnt[s[right]] == 0 {
				differ--
			}
		}

		for differ == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}
			if _, ok := dCnt[s[left]]; ok {
				dCnt[s[left]]++
				if dCnt[s[left]] == 1 {
					differ++
				}
			}
			left++
		}
	}
	return
}
