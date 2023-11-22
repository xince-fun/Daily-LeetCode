package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var pre, ans int

func getMinimumDifference(root *TreeNode) int {
	ans, pre = 500000, -1
	travel(root)
	return ans
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}
	travel(root.Left)
	if pre != -1 {
		ans = min(ans, root.Val-pre)
	}
	pre = root.Val
	travel(root.Right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
