package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans int

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 0
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left), dfs(root.Right)
	ans += abs(left - right)
	return left + right + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
