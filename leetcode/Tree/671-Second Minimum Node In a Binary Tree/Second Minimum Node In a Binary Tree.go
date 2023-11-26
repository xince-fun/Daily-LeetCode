package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans int

func findSecondMinimumValue(root *TreeNode) int {
	ans = -1
	dfs(root, root.Val)
	return ans
}

func dfs(root *TreeNode, cur int) {
	if root == nil {
		return
	}
	if root.Val != cur {
		if ans == -1 {
			ans = root.Val
		} else {
			ans = min(ans, root.Val)
		}
		return
	}
	dfs(root.Left, cur)
	dfs(root.Right, cur)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
