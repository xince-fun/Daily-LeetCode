package main

import (
	"strconv"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func binaryTreePaths(root *TreeNode) (ans []string) {
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path = path + strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path)
			return
		}
		path = path + "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return
}
