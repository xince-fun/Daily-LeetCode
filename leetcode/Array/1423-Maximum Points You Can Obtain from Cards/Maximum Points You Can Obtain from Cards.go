package main

import (
	"fmt"
	"math"
)

// 最小化剩下的牌总和 n - k
func maxScore(cardPoints []int, k int) int {
	ans := 0
	sum, minSum := 0, math.MaxInt32
	left, n := 0, len(cardPoints)
	for right := 0; right < n; right++ {
		sum += cardPoints[right]
		for right-left > k {
			sum -= cardPoints[left]
			left++
		}
		minSum = min(minSum, sum)
	}
	for i := 0; i < n; i++ {
		ans += cardPoints[i]
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
