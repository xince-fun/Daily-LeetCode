package main

// 求最长 另一个单调栈模版
func longestWPI(hours []int) (ans int) {
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
