package main

import (
	"fmt"
	"math"
	"sort"
)

/*



 */

// func minSpeedOnTime(dist []int, hour float64) int {
// 	// [left, right]
// 	left := 1
// 	right := int(1e7)
// 	for left <= right {
// 		// 循环不变量：
// 		// left-1 的回答一定是「否」
// 		// right 的回答一定是「是」
// 		mid := (left + right) / 2

// 	}
// }

func minSpeedOnTime1(dist []int, hour float64) int {
	// [left, right)
	n := len(dist)
	res := sort.Search(1e9, func(x int) bool {
		x++
		cnt := float64(0)
		for i, d := range dist {
			if i < n-1 {
				cnt += float64((d + x - 1) / x)
			} else {
				cnt += float64(d) / float64(x)
			}
		}
		return cnt <= hour
	})
	if res == 1e9 {
		return -1
	}
	return res + 1
}

func minSpeedOnTime(dist []int, hour float64) int {
	h100 := int(math.Round(hour * 100))
	n := len(dist)
	if h100 <= (n-1)*100 {
		return -1
	}
	return 1 + sort.Search(1e7-1, func(v int) bool {
		v++
		h := n - 1
		for _, d := range dist[:n-1] {
			h += (d - 1) / v
		}
		return (h*v+dist[n-1])*100 <= h100*v
	})
}

func main() {
	dist := []int{1, 3, 2}
	hour := float64(1.9)
	fmt.Println(minSpeedOnTime(dist, hour))
}
