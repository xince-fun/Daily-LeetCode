package main

func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}
	j := 0
	for _, b := range str1 {
		c := byte(b) + 1
		if b == 'z' {
			c = 'a'
		}
		if byte(b) == str2[j] || c == str2[j] {
			j++
			if j == len(str2) {
				return true
			}
		}
	}
	return false
}
