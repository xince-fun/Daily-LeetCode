package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}

func isSym(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val {
		return false
	}

	outside := isSym(left.Left, right.Right)
	inside := isSym(left.Right, right.Left)

	return outside && inside
}
