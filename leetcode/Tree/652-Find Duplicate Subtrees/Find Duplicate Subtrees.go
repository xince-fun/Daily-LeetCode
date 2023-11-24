package main

import (
	"strconv"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var ans []*TreeNode
var m map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans = make([]*TreeNode, 0)
	m = make(map[string]int)
	dfs(root)
	return ans
}

func dfs(root *TreeNode) string {
	if root == nil {
		return ""
	}
	key := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
	if m[key] == 1 {
		ans = append(ans, root)
	}
	m[key]++
	return key
}
