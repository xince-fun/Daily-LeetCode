package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var a, b, last *TreeNode

func recoverTree(root *TreeNode) {
	a, b, last = nil, nil, nil
	dfs(root)
	a.Val, b.Val = b.Val, a.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if last != nil && last.Val > root.Val {
		if a == nil {
			a = last
			b = root
		} else {
			b = root
		}
	}
	last = root
	dfs(root.Right)
}
