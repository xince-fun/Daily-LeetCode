package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans, target int

func pathSum(root *TreeNode, targetSum int) int {
	ans, target = 0, targetSum
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs2(root, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

func dfs2(root *TreeNode, val int) {
	if val == target {
		ans++
	}
	if root.Left != nil {
		dfs2(root.Left, val+root.Left.Val)
	}
	if root.Right != nil {
		dfs2(root.Right, val+root.Right.Val)
	}
}
