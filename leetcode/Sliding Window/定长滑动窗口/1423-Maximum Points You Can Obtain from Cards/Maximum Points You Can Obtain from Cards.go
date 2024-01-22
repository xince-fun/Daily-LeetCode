package main

import (
	"fmt"
)

// 最小化剩下的牌总和 n - k
func maxScore(cardPoints []int, k int) int {
	ans := 0
	sum := 0
	n := len(cardPoints)
	m := n - k
	for _, p := range cardPoints[:m] {
		sum += p
	}
	minSum := sum
	for i := m; i < n; i++ {
		sum += cardPoints[i] - cardPoints[i-m]
		minSum = min(minSum, sum)
	}
	for _, p := range cardPoints {
		ans += p
	}
	return ans - minSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Println(maxScore(cardPoints, k))
}
