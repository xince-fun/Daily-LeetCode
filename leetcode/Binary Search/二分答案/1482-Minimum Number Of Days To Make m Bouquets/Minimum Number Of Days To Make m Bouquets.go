package main

import (
	"fmt"
	"sort"
)

func minDays1(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n { // 严格不对
		return -1
	}
	check := func(d int) bool {
		f := make([]bool, n)
		for i := 0; i < n; i++ {
			if d >= bloomDay[i] {
				f[i] = true
			}
		}
		cnt := 0
		for i := 0; i < n && cnt < m; {
			if f[i] {
				cur, j := 1, i
				for cur < k && j+1 < n && f[j+1] {
					j++
					cur++
				}
				if cur == k {
					cnt++
				}
				i = j + 1
			} else {
				i++
			}
		}
		return cnt >= m
	}
	// 区间 [left, right]
	left := 1
	right := int(1e9)
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「否」
		// right+1 的回答一定是「是」
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func minDays2(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n { // 严格不对
		return -1
	}
	check := func(d int) bool {
		cnt := 0
		i := 0
		for i < n {
			if d >= bloomDay[i] {
				x := 0
				for i < n {
					if bloomDay[i] > d {
						break
					}
					i++
					x++
				}
				cnt += x / k
			} else {
				i++
			}
		}
		return cnt >= m
	}
	// 区间 [left, right]
	left := 1
	right := int(1e9)
	for left <= right {
		// 循环不变量：
		// left-1 的回答一定是「否」
		// right+1 的回答一定是「是」
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n { // 严格不对
		return -1
	}
	return sort.Search(1e9, func(d int) bool {
		cnt := 0
		i := 0
		for i < n {
			if d >= bloomDay[i] {
				x := 0
				for i < n {
					if bloomDay[i] > d {
						break
					}
					i++
					x++
				}
				cnt += x / k
			} else {
				i++
			}
		}
		return cnt >= m
	})
}

func main() {
	bloomDay := []int{7, 7, 7, 7, 12, 7, 7}
	m := 2
	k := 3
	fmt.Println(minDays(bloomDay, m, k))
}
