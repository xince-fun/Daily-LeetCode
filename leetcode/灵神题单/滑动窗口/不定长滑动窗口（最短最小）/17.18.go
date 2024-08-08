package main

func shortestSeq(big []int, small []int) (ans []int) {
	cntB, cntS := map[int]int{}, map[int]int{}
	ansLeft, ansRight, left, less := -1, len(big), 0, 0
	for _, x := range small {
		if cntS[x] == 0 {
			less++
		}
		cntS[x]++
	}
	for right, x := range big {
		cntB[x]++
		if cntB[x] == cntS[x] {
			less--
		}
		for less == 0 {
			if right-left < ansRight-ansLeft {
				ansLeft, ansRight = left, right
			}
			out := big[left]
			if cntB[out] == cntS[out] {
				less++
			}
			cntB[out]--
			left++
		}
	}
	if ansLeft < 0 {
		return
	}
	return []int{ansLeft, ansRight}
}
