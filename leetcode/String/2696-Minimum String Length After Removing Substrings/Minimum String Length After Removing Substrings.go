package main

import (
	"fmt"
	"strings"
)

func minLength(s string) int {
	str1, str2 := "AB", "CD"
	for {
		index1, index2 := strings.Index(s, str1), strings.Index(s, str2)
		if index1 != -1 {
			s = s[:index1] + s[index1+2:]
		} else if index2 != -1 {
			s = s[:index2] + s[index2+2:]
		}
		if index1 == -1 && index2 == -1 {
			break
		}
	}
	return len(s)
}

func main() {
	s := "ACBBD"
	fmt.Println(minLength(s))
}
