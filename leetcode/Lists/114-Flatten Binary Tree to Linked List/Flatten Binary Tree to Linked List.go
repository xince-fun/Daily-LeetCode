package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func flatten1(root *TreeNode) {
	st := make([]*TreeNode, 0)
	if root == nil {
		return
	}
	st = append(st, root)
	res := make([]*TreeNode, 0)
	for len(st) != 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		res = append(res, node)
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
	for i := 1; i < len(res); i++ {
		prev, cur := res[i-1], res[i]
		prev.Left, prev.Right = nil, cur
	}
}

func preorder(root *TreeNode) (ans []*TreeNode) {
	if root != nil {
		ans = append(ans, root)
		ans = append(ans, preorder(root.Left)...)
		ans = append(ans, preorder(root.Right)...)
	}
	return
}

func flatten(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}
