package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	// 单调栈
	st := []int{}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x := maxHeights[i]
		for len(st) > 0 && x < maxHeights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = maxHeights[st[len(st)-1]]
		}
		st = append(st, i)
	}
	fmt.Println(right)
	return 0
}

func main() {
	maxHeights := []int{5, 3, 4, 1, 1}
	maximumSumOfHeights(maxHeights)
}
