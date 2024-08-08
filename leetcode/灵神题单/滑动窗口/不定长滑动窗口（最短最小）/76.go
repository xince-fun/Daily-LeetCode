package main

func minWindow1(s, t string) (ans string) {
	if len(t) > len(s) {
		return ""
	}

	cntT := map[rune]int{}
	for _, c := range t {
		cntT[c]++
	}
	diff := len(cntT)
	minLen := len(s) + 1

	left := 0
	for right, c := range s {
		if _, ok := cntT[c]; ok {
			cntT[c]--
			if cntT[c] == 0 {
				diff--
			}
		}

		for diff == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}
			if _, ok := cntT[rune(s[left])]; ok {
				cntT[rune(s[left])]++
				if cntT[rune(s[left])] == 1 {
					diff++
				}
			}
			left++
		}
	}
	return
}

func isCover(cntS, cntT [128]int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow2(s, t string) (ans string) {
	m := len(s)
	ansLeft, ansRight, left := -1, m, 0
	var cntS, cntT [128]int
	for _, c := range t {
		cntT[c]++
	}

	for right, c := range s {
		cntS[c]++
		for isCover(cntS, cntT) {
			if right-left < ansRight-ansLeft {
				ansLeft, ansRight = left, right
			}
			cntS[s[left]]--
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

func minWindow(s, t string) (ans string) {
	m := len(s)
	ansLeft, ansRight, left := -1, m, 0
	less := 0
	var cntS, cntT [128]int
	for _, c := range t {
		if cntT[c] == 0 {
			less++
		}
		cntT[c]++
	}
	for right, c := range s {
		cntS[c]++
		if cntS[c] == cntT[c] {
			less--
		}
		for less == 0 {
			if right-left < ansRight-ansLeft {
				ansLeft, ansRight = left, right
			}
			if cntS[s[left]] == cntT[s[left]] {
				less++
			}
			cntS[s[left]]--
			left++
		}
	}
	if ansLeft < 0 {
		return
	}
	return s[ansLeft : ansRight+1]
}
