package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	level := 2
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if level == depth {
				left, right := node.Left, node.Right
				node.Left = &TreeNode{Val: val, Left: left}
				node.Right = &TreeNode{Val: val, Right: right}
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if level == depth {
			return root
		}
		level++
	}
	return root
}
