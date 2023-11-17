package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		len := len(q)
		res := []int{}
		for i := 0; i < len; i++ {
			cur := q[0]
			q = q[1:]
			res = append(res, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		ans = append(ans, res)
	}
	return ans
}
