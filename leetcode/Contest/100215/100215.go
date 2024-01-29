package main

import "fmt"

func countKeyChanges(s string) (ans int) {
	c := s[0]
	for i := 1; i < len(s); i++ {
		if c != s[i] && c != (s[i]^32) {
			ans++
		}
		c = s[i]
	}
	return
}

func main() {
	s := "AaAaAaaA"
	fmt.Println(countKeyChanges(s))
}
