package main

func winnerOfGame(colors string) bool {
	cnt1, cnt2 := 0, 0
	for i, n := 0, len(colors); i < n; {
		i0 := i
		for i++; i < n && colors[i] == colors[i0]; i++ {
		}
		if i-i0 >= 3 {
			if colors[i0] == 'A' {
				cnt1++
			} else {
				cnt2++
			}
		}
	}
	return cnt1 > cnt2
}
