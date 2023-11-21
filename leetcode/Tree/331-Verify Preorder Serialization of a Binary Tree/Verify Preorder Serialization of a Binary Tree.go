package main

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	arr := strings.Split(preorder, ",")
	in, out := 0, 0
	n := len(arr)
	for i, a := range arr {
		if a != "#" {
			out += 2
		}
		if i != 0 {
			in++
		}
		if i != n-1 && out <= in {
			return false
		}
	}
	return in == out
}

func main() {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"

	isValidSerialization(str)
}
