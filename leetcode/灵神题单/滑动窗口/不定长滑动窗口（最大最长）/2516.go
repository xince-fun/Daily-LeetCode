package main

func takeCharacters(s string, k int) int {
	cntA, cntB, cntC := 0, 0, 0
	for _, c := range s {
		if c == 'a' {
			cntA++
		} else if c == 'b' {
			cntB++
		} else {
			cntC++
		}
	}
	if cntA < k || cntB < k || cntC < k {
		return -1
	}
	leftA, leftB, leftC := cntA-k, cntB-k, cntC-k
	cntA, cntB, cntC = 0, 0, 0
	left := 0
	ans := 0
	for right, c := range s {
		if c == 'a' {
			cntA++
		} else if c == 'b' {
			cntB++
		} else {
			cntC++
		}

		for cntA > leftA || cntB > leftB || cntC > leftC {
			out := s[left]
			left++
			if out == 'a' {
				cntA--
			} else if out == 'b' {
				cntB--
			} else {
				cntC--
			}
		}
		ans = max(ans, right-left+1)
	}
	return len(s) - ans
}
