package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans []int
var base, count, maxCount int

func findMode(root *TreeNode) []int {
	ans = make([]int, 0)
	base, count, maxCount = 0, 0, 0
	travel(root)
	return ans
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}
	travel(root.Left)
	update(root.Val)
	travel(root.Right)
}

func update(x int) {
	if x == base {
		count++
	} else {
		count = 1
		base = x
	}
	if count == maxCount {
		ans = append(ans, base)
	}
	if count > maxCount {
		maxCount = count
		ans = []int{base}
	}
}
