package main

// 前缀和 + 单调栈
func longestWPI1(hours []int) (ans int) {
	n := len(hours)
	s := make([]int, n+1)
	st := []int{0} // s[0]
	for j, h := range hours {
		j++
		s[j] = s[j-1]
		if h > 8 {
			s[j]++
		} else {
			s[j]--
		}
		if s[j] < s[st[len(st)-1]] {
			st = append(st, j)
		}
	}
	for i := n; i > 0; i-- {
		for len(st) > 0 && s[i] > s[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 前缀和
func longestWPI(hours []int) (ans int) {
	s := 0
	pos := make([]int, len(hours)+2)
	for i, h := range hours {
		i++
		if h > 8 {
			s--
		} else {
			s++
		}
		if s < 0 {
			ans = i
		} else {
			if pos[s+1] > 0 {
				ans = max(ans, i-pos[s+1])
			}
			if pos[s] == 0 {
				pos[s] = i
			}
		}
	}
	return
}
