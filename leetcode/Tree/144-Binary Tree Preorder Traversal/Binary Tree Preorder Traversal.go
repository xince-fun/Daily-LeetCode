package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func preorderTraversal1(root *TreeNode) []int {
	nums := []int{}
	travel(root, &nums)
	return nums
}

func travel(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	travel(root.Left, nums)
	travel(root.Right, nums)
}

func preorderTraversal(root *TreeNode) []int {
	nums := []int{}
	st := make([]*TreeNode, 0)
	if root == nil {
		return nums
	}
	st = append(st, root)
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		nums = append(nums, node.Val)
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}
	return nums
}
