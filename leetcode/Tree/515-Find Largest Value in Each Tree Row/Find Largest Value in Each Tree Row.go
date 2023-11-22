package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func largestValues1(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		maxNum := q[0].Val
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Val > maxNum {
				maxNum = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, maxNum)
	}
	return ans
}

func largestValues(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, currentHeight int) {
		if node == nil {
			return
		}
		if currentHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[currentHeight] = max(ans[currentHeight], node.Val)
		}
		dfs(node.Left, currentHeight+1)
		dfs(node.Right, currentHeight+1)
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
