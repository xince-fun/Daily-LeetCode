package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	res := []*TreeNode{}
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for mid := lo; mid <= hi; mid++ {
		leftTree := build(lo, mid-1)
		rightTree := build(mid+1, hi)
		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{Val: mid}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
