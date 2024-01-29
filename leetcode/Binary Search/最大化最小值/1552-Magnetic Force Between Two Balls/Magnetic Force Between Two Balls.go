package main

import (
	"fmt"
	"sort"
)

// 最大化最小值
func maxDistance1(position []int, m int) int {
	sort.Ints(position)
	left, right := 0, (position[len(position)-1]-position[0])/(m-1)+1

	check := func(mid int) bool {
		x0 := position[0]
		cnt := 1
		for _, x := range position {
			if x >= x0+mid {
				x0 = x
				cnt += 1
			}
		}
		return cnt >= m
	}

	// 开区间
	for left+1 < right {
		mid := (left + right + 1) / 2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	right := int(1e9)
	return sort.Search(right, func(mid int) bool {
		x0 := position[0]
		cnt := 1
		for _, x := range position {
			if x >= x0+mid {
				x0 = x
				cnt += 1
			}
		}
		return cnt < m
	}) - 1
}

func main() {
	position := []int{5, 4, 3, 2, 1, 1000000000}
	m := 2
	fmt.Println(maxDistance(position, m))
}
