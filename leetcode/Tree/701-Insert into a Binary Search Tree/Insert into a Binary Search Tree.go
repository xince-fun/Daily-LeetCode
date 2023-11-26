package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur, parent := root, root
	for cur != nil {
		parent = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	node := &TreeNode{Val: val}
	if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return root
}
