package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func reverseOddLevels1(root *TreeNode) *TreeNode {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	isOdd := false
	for len(q) > 0 {
		length := len(q)
		st := []*TreeNode{}
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			st = append(st, node)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if isOdd {
			for i, j := 0, len(st)-1; i < j; i, j = i+1, j-1 {
				st[i].Val, st[j].Val = st[j].Val, st[i].Val
			}
		}
		isOdd = !isOdd
	}
	return root
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, true)
	return root
}

func dfs(left, right *TreeNode, isOdd bool) {
	if left == nil {
		return
	}
	if isOdd {
		left.Val, right.Val = right.Val, left.Val
	}
	dfs(left.Left, right.Right, !isOdd)
	dfs(left.Right, right.Left, !isOdd)
}
