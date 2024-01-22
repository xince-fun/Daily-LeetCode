package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	maxIdx := len(s) - 1
	p, q := -1, 0
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			p, q = i, maxIdx
		}
	}
	if p == -1 {
		return num
	}
	t := []byte(s)
	t[p], t[q] = t[q], t[p]
	ans, _ := strconv.Atoi(string(t))
	return ans
}

func main() {
	num := 2736
	fmt.Println(maximumSwap(num))
}
