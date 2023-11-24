package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func averageOfLevels(root *TreeNode) (ans []float64) {
	var dfs func(*TreeNode, int)
	count := []int{}
	dfs = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, float64(node.Val))
			count = append(count, 1)
		} else {
			ans[curHeight] = (ans[curHeight]*float64(count[curHeight]) + float64(node.Val)) / float64(count[curHeight]+1)
			count[curHeight]++
		}
		dfs(node.Left, curHeight+1)
		dfs(node.Right, curHeight+1)
	}
	dfs(root, 0)
	return
}
