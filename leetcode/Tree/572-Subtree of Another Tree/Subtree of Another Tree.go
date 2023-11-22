package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if check(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return check(root1.Left, root2.Left) && check(root1.Right, root2.Right)
}
