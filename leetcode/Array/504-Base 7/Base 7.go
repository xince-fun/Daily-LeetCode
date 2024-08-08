package main

import "strconv"

func convertToBase7(num int) string {
	ss := ""
	if num < 0 {
		ss += "-"
		num = -num
	}
	for num > 0 {
		s := strconv.Itoa(num % 7)
		ss += s
		num /= 7
	}
	return reverse(ss)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
