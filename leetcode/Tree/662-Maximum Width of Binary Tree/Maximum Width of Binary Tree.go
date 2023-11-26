package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func widthOfBinaryTree(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	q, res := []*TreeNode{}, 0
	q = append(q, &TreeNode{Val: 0, Left: root.Left, Right: root.Right})

	for len(q) > 0 {
		var left, right *int
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				newVal := node.Val * 2
				q = append(q, &TreeNode{Val: newVal, Left: node.Left.Left, Right: node.Left.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
			if node.Right != nil {
				newVal := node.Val*2 + 1
				q = append(q, &TreeNode{Val: newVal, Left: node.Right.Left, Right: node.Right.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
		}
		switch {
		case left != nil && right == nil || left == nil && right != nil:
			res = max(res, 1)
		case left != nil && right != nil:
			res = max(res, *right-*left+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println(widthOfBinaryTree(root))
}
