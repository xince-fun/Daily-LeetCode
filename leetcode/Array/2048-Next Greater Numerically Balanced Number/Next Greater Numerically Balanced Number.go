package main

import (
	"fmt"
)

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		var nums [10]int
		for y := i; y > 0; y /= 10 {
			nums[y%10]++
		}
		isTrue := true
		for y := i; y > 0; y /= 10 {
			if y%10 != nums[y%10] {
				isTrue = false
				break
			}
		}
		if isTrue {
			return i
		}
	}
}

func main() {
	n := 1
	fmt.Println(nextBeautifulNumber(n))
}
