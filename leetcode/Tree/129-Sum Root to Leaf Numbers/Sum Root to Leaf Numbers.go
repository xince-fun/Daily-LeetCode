package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans int

func sumNumbers(root *TreeNode) int {
	ans = 0
	dfs(root, 0)
	return ans
}

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		ans += sum*10 + root.Val
		return
	}

	dfs(root.Left, sum*10+root.Val)
	dfs(root.Right, sum*10+root.Val)
}
