package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	x := root.Val
	if x > high {
		return rangeSumBST(root.Left, low, high)
	}
	if x < low {
		return rangeSumBST(root.Right, low, high)
	}
	return x + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
