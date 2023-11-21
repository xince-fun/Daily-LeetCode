package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfs(root.Left)
	r0, r1 := dfs(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
