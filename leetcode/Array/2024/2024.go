package main

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	left := 0
	cntT, cntF := 0, 0
	for right, c := range answerKey {
		if c == 'T' {
			cntT++
		} else {
			cntF++
		}

		minA := min(cntF, cntT)
		for minA > k {
			out := answerKey[left]
			if out == 'T' {
				cntT--
			} else {
				cntF--
			}
			minA = min(cntF, cntT)
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
