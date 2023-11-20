package main

import (
	"math"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func countNodes(root *TreeNode) int {
	left, right := root, root
	leftH, rightH := 0, 0
	for left != nil {
		left = left.Left
		leftH++
	}
	for right != nil {
		right = right.Right
		rightH++
	}
	if leftH == rightH {
		return int(math.Pow(2, float64(leftH)) - 1)
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
