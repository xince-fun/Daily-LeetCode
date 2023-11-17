package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	return build(preorder, inorder, 0, n-1, 0, n-1)
}

func build(preorder []int, inorder []int, preL, preR, inL, inR int) *TreeNode {
	if preL > preR {
		return nil
	}

	root := &TreeNode{Val: preorder[preL]}

	index := 0
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == preorder[preL] {
			index = i
			break
		}
	}

	numLeft := index - inL
	root.Left = build(preorder, inorder, preL+1, preL+numLeft, inL, index-1)
	root.Right = build(preorder, inorder, preL+numLeft+1, preR, index+1, inR)
	return root
}
