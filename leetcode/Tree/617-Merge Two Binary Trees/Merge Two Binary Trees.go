package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 != nil && root2 == nil {
		return &TreeNode{Val: root1.Val}
	}
	if root1 == nil && root2 != nil {
		return &TreeNode{Val: root2.Val}
	}
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
