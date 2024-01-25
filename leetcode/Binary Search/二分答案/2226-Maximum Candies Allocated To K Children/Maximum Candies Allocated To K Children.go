package main

import (
	"fmt"
	"sort"
)

/*
可以将每堆糖果分成任意数量的子堆
candies = [5, 8, 6], k = 3
=> 5 5 3 5 1

candies = [2, 5] k = 11

如果能分 每个人最少分到1，最多分到 1e9
下界为 1

*/

func maximumCandies1(candies []int, k int64) int {
	// 上界 sum(candies) / k
	right := int64(0)
	for _, c := range candies {
		right += int64(c)
	}
	right /= k
	// [left, right]
	left := int64(1)
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「是」
		// right+1 的回答一定为「否」
		mid := (left + right) / 2
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c) / mid
		}
		// 如果大于等于k，说明满足
		if cnt >= k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int(right)
}

func maximumCandies2(candies []int, k int64) int {
	// (left, right]
	right := int64(0)
	for _, c := range candies {
		right += int64(c)
	}
	right /= k
	left := int64(0)
	for left < right {
		// 循环不变量：
		// left 的回答一定是「是」
		// right+1 的回答一定为「否」
		mid := (left + right + 1) / 2
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c) / mid
		}
		// 如果大于等于k，说明满足
		if cnt >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return int(left)
}

func maximumCandies3(candies []int, k int64) int {
	// [left, right)
	right := int64(0)
	for _, c := range candies {
		right += int64(c)
	}
	right = right/k + 1
	left := int64(1)
	for left < right {
		// 循环不变量：
		// left 的回答一定是「是」
		// right+1 的回答一定为「否」
		mid := (left + right) / 2
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c) / mid
		}
		// 如果大于等于k，说明满足
		if cnt >= k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int(left) - 1
}

func maximumCandies4(candies []int, k int64) int {
	// (left, right)
	right := int64(0)
	for _, c := range candies {
		right += int64(c)
	}
	right = right/k + 1
	left := int64(0)
	for left+1 < right {
		// 循环不变量：
		// left 的回答一定是「是」
		// right 的回答一定为「否」
		mid := (left + right + 1) / 2
		cnt := int64(0)
		for _, c := range candies {
			cnt += int64(c) / mid
		}
		// 如果大于等于k，说明满足
		if cnt >= k {
			left = mid
		} else {
			right = mid
		}
	}
	return int(left)
}

func maximumCandies(candies []int, k int64) int {
	// [left, right)
	right := 0
	for _, c := range candies {
		right += c
	}
	right = right/int(k) + 1
	return sort.Search(right, func(i int) bool {
		i++
		cnt := int64(0)
		for _, candy := range candies {
			cnt += int64(candy / i)
		}
		return cnt < k
	})
}

func main() {
	candies := []int{5, 8, 6}
	k := 3
	fmt.Println(maximumCandies(candies, int64(k)))
}
