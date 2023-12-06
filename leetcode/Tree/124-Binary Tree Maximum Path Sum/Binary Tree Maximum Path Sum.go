package main

import (
	"math"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var res int

func maxPathSum1(root *TreeNode) int {
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

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := dfs(root.Left), dfs(root.Right)
		ans = max(ans, left+right+root.Val)
		return max(max(left, right)+root.Val, 0)
	}
	dfs(root)
	return ans
}
