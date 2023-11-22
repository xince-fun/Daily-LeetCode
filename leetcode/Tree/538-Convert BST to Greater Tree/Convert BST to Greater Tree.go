package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

// var sum int

// func convertBST1(root *TreeNode) *TreeNode {
// 	sum = 0
// 	bfs(root)
// 	bfs2(root)
// 	return root
// }

// func bfs(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	bfs(root.Left)
// 	sum += root.Val
// 	bfs(root.Right)
// }

// func bfs2(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	bfs2(root.Left)
// 	tmp := root.Val
// 	root.Val = sum
// 	sum -= tmp
// 	bfs2(root.Right)
// }

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	tmp := root.Val
	root.Val += sum
	sum += tmp
	dfs(root.Left)
}
