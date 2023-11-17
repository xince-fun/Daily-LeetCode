package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var index = make(map[int]int)

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	for i := 0; i < n; i++ {
		index[inorder[i]] = i
	}
	return build(inorder, postorder, 0, n-1, 0, n-1)
}

func build(inorder, postorder []int, inL, inR, posL, posR int) *TreeNode {
	if posL > posR {
		return nil
	}

	index_root := index[postorder[posR]]

	root := &TreeNode{Val: postorder[posR]}

	numLeft := index_root - inL
	root.Left = build(inorder, postorder, inL, index_root-1, posL, posL+numLeft-1)
	root.Right = build(inorder, postorder, index_root+1, inR, posL+numLeft, posR-1)
	return root
}
