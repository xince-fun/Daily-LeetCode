package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		// 保证 p.Val <= q.Val
		return lowestCommonAncestor(root, q, p)
	}
	if root.Val >= p.Val && root.Val <= q.Val {
		// p <= root <= q
		return root
	}
	if root.Val > q.Val {
		// p 和 q 都在左子树
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
