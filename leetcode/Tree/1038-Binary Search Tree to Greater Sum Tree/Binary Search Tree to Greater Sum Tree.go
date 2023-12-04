package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	tmp := root.Val
	root.Val += sum
	sum += tmp
	dfs(root.Left)
}
