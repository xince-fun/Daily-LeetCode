package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	travel(root, &ans)
	return ans
}

func travel(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	travel(root.Left, ans)
	travel(root.Right, ans)
	*ans = append(*ans, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	st := []*TreeNode{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			st = st[:len(st)-1]
			st = append(st, node)
			st = append(st, nil)
			if node.Right != nil {
				st = append(st, node.Right)
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			st = st[:len(st)-1]
			node := st[len(st)-1]
			st = st[:len(st)-1]
			ans = append(ans, node.Val)
		}
	}
	return ans
}
