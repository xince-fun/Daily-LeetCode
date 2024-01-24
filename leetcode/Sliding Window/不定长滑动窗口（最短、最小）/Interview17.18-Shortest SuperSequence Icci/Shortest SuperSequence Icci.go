package main

func shortestSeq1(big []int, small []int) (ans []int) {
	if len(small) > len(big) {
		return
	}

	dCnt := make(map[int]int)
	for i := 0; i < len(small); i++ {
		dCnt[small[i]]++
	}

	cCnt := make(map[int]int)
	differ := len(small)

	left := 0
	minLen := len(big) + 1

	for right := 0; right < len(big); right++ {
		if dCnt[big[right]] == 0 {
			continue
		}

		cCnt[big[right]]++
		if cCnt[big[right]] <= dCnt[big[right]] {
			differ--
		}

		for differ == 0 {
			if dCnt[big[left]] == 0 {
				left++
				continue
			}

			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = []int{left, right}
			}
			cCnt[big[left]]--
			if cCnt[big[left]] < dCnt[big[left]] {
				differ++
			}
			left++
		}
	}
	return
}

func shortestSeq(big []int, small []int) (ans []int) {
	if len(small) > len(big) {
		return
	}

	dCnt := make(map[int]int)
	for i := 0; i < len(small); i++ {
		dCnt[small[i]]++
	}

	// 这里写 len(small)和 len(dCnt) 效果是一样的 因为题目说了元素不相同
	// 但是如果有重复，则需要写 len(dCnt) 所以最好都写 len(dCnt)
	differ := len(dCnt)
	left := 0
	minLen := len(big) + 1

	for right := 0; right < len(big); right++ {
		if _, ok := dCnt[big[right]]; ok {
			dCnt[big[right]]--
			if dCnt[big[right]] == 0 {
				differ--
			}
		}

		for differ == 0 {

			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = []int{left, right}
			}
			if _, ok := dCnt[big[left]]; ok {
				dCnt[big[left]]++
				if dCnt[big[left]] == 1 {
					differ++
				}
			}
			left++
		}
	}
	return
}
