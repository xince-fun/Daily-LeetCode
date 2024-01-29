package main

// 分组循环
func maxPower(s string) (ans int) {
	n := len(s)
	i := 0
	for i < n {
		start := i
		c := s[start]
		i++
		for i < n && s[i] == c {
			i++
		}
		ans = max(ans, i-start)
	}
	return
}
