package main

import (
	"math"
	"strconv"
)

// 背包容量为 9
//完全背包最多可以选择的数字
//f[i][j] = max(f[i-1][j], f[i][j-cost[i-1]]+1)

func largestNumber(cost []int, target int) string {
	ms := ""
	n := len(cost)
	f := make([]int, target+1)
	for i := range f {
		f[i] = math.MinInt
	}
	f[0] = 0
	for i := 0; i < n; i++ {
		c := cost[i]
		for j := c; j <= target; j++ {
			f[j] = max(f[j], f[j-c]+1)
		}
	}
	if f[target] <= 0 {
		return "0"
	}
	for i, j := 9, target; i >= 1; i-- {
		c := cost[i-1]
		for j-c >= 0 && f[j] == f[j-c]+1 {
			ms += strconv.Itoa(i)
			j = j - c
		}
	}
	return ms
}
