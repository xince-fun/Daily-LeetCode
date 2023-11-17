package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return check(root.Left, min, root) && check(root.Right, root, max)
}
