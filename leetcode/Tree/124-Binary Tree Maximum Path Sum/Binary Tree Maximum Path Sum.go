package main

import (
	"math"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	if root == nil {
		return 0
	}
	oneSideMax(root)
	return res
}

func oneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxSum := max(0, oneSideMax(root.Left))
	rightMaxSum := max(0, oneSideMax(root.Right))

	pathMaxSum := root.Val + leftMaxSum + rightMaxSum
	res = max(res, pathMaxSum)
	return max(leftMaxSum, rightMaxSum) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
