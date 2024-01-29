package main

// 分组循环
func checkZeroOnes(s string) bool {
	m1, m2 := 0, 0
	n := len(s)
	i := 0
	for i < n {
		start := i
		v := s[start]
		for i < n && s[i] == v {
			i++
		}
		if v == '1' {
			m1 = max(m1, i-start)
		} else {
			m2 = max(m2, i-start)
		}
	}
	return m1 > m2
}
