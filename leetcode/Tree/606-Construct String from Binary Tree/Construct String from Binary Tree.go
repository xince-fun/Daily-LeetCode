package main

import (
	"strconv"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var str string

func tree2str(root *TreeNode) string {
	str = ""
	preorder(root)
	return str[1 : len(str)-1]
}

func preorder(root *TreeNode) {
	str += "("
	str += strconv.Itoa(root.Val)
	if root.Left != nil {
		preorder(root.Left)
	} else if root.Right != nil {
		str += "()"
	}
	if root.Right != nil {
		preorder(root.Right)
	}
	str += ")"
}
