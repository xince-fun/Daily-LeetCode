package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	return count(1, n)
}

func count(lo, hi int) int {
	if lo > hi {
		return 1
	}
	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}
	res := 0

	for mid := lo; mid <= hi; mid++ {
		left := count(lo, mid-1)
		right := count(mid+1, hi)
		res += left * right
	}
	memo[lo][hi] = res
	return res
}
