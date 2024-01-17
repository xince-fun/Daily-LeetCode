package main

import (
	"fmt"
)

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	INF := int(1e9)
	ans := 0
	for mask := 0; mask < (1 << n); mask++ {
		c := make(map[int]bool)
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				c[i] = true
			}
		}
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dp[i][j] = INF
			}
		}
		for _, i := range roads {
			if c[i[0]] && c[i[1]] {
				v := min(i[2], dp[i[0]][i[1]])
				dp[i[0]][i[1]] = v
				dp[i[1]][i[0]] = v
			}
		}
		for i := 0; i < n; i++ {
			dp[i][i] = 0
		}
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
				}
			}
		}
		ok := true
		for i := range c {
			for j := range c {
				if dp[i][j] > maxDistance {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	n := 3
	maxDistance := 5
	roads := [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}
	fmt.Println(numberOfSets(n, maxDistance, roads)) // Output: 5

	n = 3
	maxDistance = 5
	roads = [][]int{{0, 1, 20}, {0, 1, 10}, {1, 2, 2}, {0, 2, 2}}
	fmt.Println(numberOfSets(n, maxDistance, roads)) // Output: 7

	n = 1
	maxDistance = 10
	roads = [][]int{}
	fmt.Println(numberOfSets(n, maxDistance, roads)) // Output: 2
}
