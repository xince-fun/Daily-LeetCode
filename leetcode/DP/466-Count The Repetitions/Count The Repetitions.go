package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	count, p2 := 0, 0
	len1, len2 := len(s1), len(s2)
	for i := 0; i < n1; i++ {
		for j := 0; j < len1; j++ {
			if s1[j] == s2[p2] {
				p2++
				if p2 == len2 {
					count++
					p2 = 0
				}
			}
		}
	}
	return count / n2
}
