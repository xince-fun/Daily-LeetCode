package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Left == nil {
		return root
	}
	return findMin(root.Left)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	result := key - root.Val

	if result < 0 {
		root.Left = deleteNode(root.Left, key)
	} else if result > 0 {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left != nil && root.Right != nil {
		root.Val = findMin(root.Right).Val
		root.Right = deleteNode(root.Right, root.Val)
	} else {
		if root.Left != nil {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
