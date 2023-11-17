package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans [][]int
var path []int

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	dfs(root, targetSum)
	return ans
}

func dfs(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
	}
	dfs(root.Left, targetSum-root.Val)
	dfs(root.Right, targetSum-root.Val)
	path = path[:len(path)-1]
}
