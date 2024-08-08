package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}

	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && node.Val == sum {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, sum-node.Val)
		dfs(node.Right, sum-node.Val)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return
}
