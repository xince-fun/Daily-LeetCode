package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cur := q[0]
			q = q[1:]
			if i == length-1 {
				ans = append(ans, cur.Val)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return ans
}
