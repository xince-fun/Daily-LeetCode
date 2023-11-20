package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

type BSTIterator struct {
	Head int
	Nums []int
}

func inorder(root *TreeNode) []int {
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
			st = append(st, nil)
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

func Constructor(root *TreeNode) BSTIterator {
	nums := inorder(root)
	return BSTIterator{
		Head: -1,
		Nums: nums,
	}
}

func (this *BSTIterator) Next() int {
	this.Head = this.Head + 1
	return this.Nums[this.Head]
}

func (this *BSTIterator) HasNext() bool {
	return this.Head < len(this.Nums)-1
}
