package main

func maxConsecutiveAnswers1(answerKey string, k int) (ans int) {
	left, cnt := 0, 0
	for right, x := range answerKey {
		if x == 'F' {
			cnt++
		}
		for cnt > k {
			if answerKey[left] == 'F' {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	cnt = 0
	for right, x := range answerKey {
		if x == 'T' {
			cnt++
		}
		for cnt > k {
			if answerKey[left] == 'T' {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	calc := func(target byte) (ans int) {
		left, cnt := 0, 0
		for right := range answerKey {
			if answerKey[right] == target {
				cnt++
			}
			for cnt > k {
				if answerKey[left] == target {
					cnt--
				}
				left++
			}
			ans = max(ans, right-left+1)
		}
		return
	}
	return max(calc('T'), calc('F'))
}
