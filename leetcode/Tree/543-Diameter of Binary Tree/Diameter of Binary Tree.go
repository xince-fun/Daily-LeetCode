package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left), dfs(root.Right)
	ans = max(ans, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
