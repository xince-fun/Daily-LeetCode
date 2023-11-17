package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	curIsLeft := false
	for len(q) > 0 {
		leng := len(q)
		res := []int{}
		st := []*TreeNode{}
		for i := 0; i < leng; i++ {
			cur := q[0]
			q = q[1:]
			res = append(res, cur.Val)
			if curIsLeft {
				if cur.Right != nil {
					st = append(st, cur.Right)
				}
				if cur.Left != nil {
					st = append(st, cur.Left)
				}
			} else {
				if cur.Left != nil {
					st = append(st, cur.Left)
				}
				if cur.Right != nil {
					st = append(st, cur.Right)
				}
			}
		}
		for len(st) > 0 {
			q = append(q, st[len(st)-1])
			st = st[:len(st)-1]
		}
		curIsLeft = !curIsLeft
		ans = append(ans, res)
	}
	return ans
}
