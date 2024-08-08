package main

func balancedString1(s string) (ans int) {
	b := len(s) / 4
	cntQ, cntW, cntE, cntR := 0, 0, 0, 0
	for _, c := range s {
		if c == 'Q' {
			cntQ++
		} else if c == 'W' {
			cntW++
		} else if c == 'E' {
			cntE++
		} else {
			cntR++
		}
	}
	if cntQ == b && cntW == b && cntE == b && cntR == b {
		return 0
	}
	ans = len(s)
	left := 0
	for right, c := range s {
		if c == 'Q' {
			cntQ--
		} else if c == 'W' {
			cntW--
		} else if c == 'E' {
			cntE--
		} else {
			cntR--
		}

		for cntQ <= b && cntW <= b && cntE <= b && cntR <= b {
			out := s[left]
			if out == 'Q' {
				cntQ++
			} else if out == 'W' {
				cntW++
			} else if out == 'E' {
				cntE++
			} else {
				cntR++
			}
			left++
		}
		ans = min(ans, right-left+1)
	}
	return
}

func balancedString(s string) (ans int) {
	cnt, m := [26]int{}, len(s)/4
	for _, c := range s {
		cnt[c-'A']++
	}
	if cnt['Q'-'A'] == m && cnt['W'-'A'] == m && cnt['E'-'A'] == m && cnt['R'-'A'] == m {
		return 0
	}

	ans = len(s)
	left := 0
	for right, c := range s {
		cnt[c-'A']--
		for cnt['Q'-'A'] <= m && cnt['W'-'A'] <= m && cnt['E'-'A'] <= m && cnt['R'-'A'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]-'A']++
			left++
		}
	}
	return
}
