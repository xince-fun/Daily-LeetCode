package main

import (
	"math/bits"
	"sort"
)

/*
func countSetBitsAtMultiplesOfX(n int64, x int) int64 {
	count := int64(0)
	for i := x; i <= 64; i += x {
		completeCycles := (n + 1) / (1 << i)
		count += completeCycles * (1 << (i - 1))

		remaining := (n + 1) % (1 << i)
		if remaining > (1 << (i - 1)) {
			count += remaining - (1 << (i - 1))
		}
	}
	return count
}

func findMaximumNumber(k int64, x int) int64 {
	left, right := int64(1), int64(2)
	for countSetBitsAtMultiplesOfX(right, x) <= k {
		right *= 2
	}

	for left < right {
		mid := (left + right) / 2
		if countSetBitsAtMultiplesOfX(mid, x) <= k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1

}
*/

/*
1. 二分 + 数位DP

给定 num， 统计 1～num 的价值和，判断价值和是否 <= k
二分相当于多花费 log 的时间， 额外增加一个条件

2. 二分 + 枚举数位

*/

func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		n := bits.Len(uint(num))
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		dfs = func(i, cnt1 int, limitHigh bool) (res int) {
			if i < 0 {
				return cnt1
			}
			if !limitHigh {
				p := &memo[i][cnt1]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}
			up := 1
			if limitHigh {
				up = num >> i & 1
			}
			for d := 0; d <= up; d++ {
				c := cnt1
				if d == 1 && (i+1)%x == 0 {
					c++
				}
				res += dfs(i-1, c, limitHigh && d == up)
			}
			return
		}
		return dfs(n-1, 0, true) > int(k)
	})
	return int64(ans)
}
