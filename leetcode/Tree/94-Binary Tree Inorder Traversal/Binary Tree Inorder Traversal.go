package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func inorderTraversal1(root *TreeNode) []int {
	nums := []int{}
	travel(root, &nums)
	return nums
}

func travel(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	travel(root.Left, nums)
	*nums = append(*nums, root.Val)
	travel(root.Right, nums)
}

func inorderTraversal(root *TreeNode) []int {
	nums := []int{}
	st := make([]*TreeNode, 0)
	if root == nil {
		return nums
	}
	st = append(st, root)
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			st = st[:len(st)-1]
			if node.Right != nil {
				st = append(st, node.Right)
			}

			st = append(st, node)
			st = append(st, nil) //
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			st = st[:len(st)-1]
			node := st[len(st)-1]
			st = st[:len(st)-1]
			nums = append(nums, node.Val)
		}
	}
	return nums
}
