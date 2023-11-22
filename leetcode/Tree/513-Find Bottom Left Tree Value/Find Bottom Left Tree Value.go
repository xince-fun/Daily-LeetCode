package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	var ans *TreeNode
	for len(q) > 0 {
		length := len(q)
		tmp := []*TreeNode{}
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = tmp[0]
	}
	return ans.Val
}
