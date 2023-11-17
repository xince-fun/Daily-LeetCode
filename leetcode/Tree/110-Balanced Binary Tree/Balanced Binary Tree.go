package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := getHeight(root.Left)
	if leftH == -1 {
		return -1 // tuichu
	}
	rightH := getHeight(root.Right)
	if rightH == -1 || abs(leftH-rightH) > 1 {
		return -1
	}
	return max(leftH, rightH) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
