package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var maxNum int

func longestUnivaluePath(root *TreeNode) int {
	maxNum = 0
	dfs(root)
	return maxNum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, cur := 0, 0
	left, right := dfs(root.Left), dfs(root.Right)

	if root.Left != nil && root.Left.Val == root.Val {
		ans = left + 1
		cur += left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		ans = max(ans, right+1)
		cur += right + 1
	}
	maxNum = max(cur, maxNum)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
