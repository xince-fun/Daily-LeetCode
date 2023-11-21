package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var path []int

func kthSmallest1(root *TreeNode, k int) int {
	path = make([]int, 0)
	travel(root)
	return path[k-1]
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}

	travel(root.Left)
	path = append(path, root.Val)
	travel(root.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	st := []*TreeNode{root}
	count := 1
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			st = st[:len(st)-1]
			if node.Right != nil {
				st = append(st, node.Right)
			}
			st = append(st, node)
			st = append(st, nil)
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			st = st[:len(st)-1]
			if count == k {
				return st[len(st)-1].Val
			}
			st = st[:len(st)-1]
			count++
		}
	}
	return -1
}
