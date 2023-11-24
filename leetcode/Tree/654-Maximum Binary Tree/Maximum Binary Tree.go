package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	num := -1
	index := left
	for i := left; i <= right; i++ {
		if nums[i] > num {
			num = nums[i]
			index = i
		}
	}
	fmt.Println(num, index)
	root := &TreeNode{Val: num}
	root.Left, root.Right = build(nums, left, index-1), build(nums, index+1, right)
	return root
}
