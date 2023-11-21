package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					ans += node.Left.Val
				}
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return ans
}
