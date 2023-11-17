package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	st := [][]int{}
	for len(q) > 0 {
		length := len(q)
		res := []int{}
		for i := 0; i < length; i++ {
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
		st = append(st, res)
	}
	for len(st) != 0 {
		ans = append(ans, st[len(st)-1])
		st = st[:len(st)-1]
	}
	return ans
}
