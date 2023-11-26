package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var path []int
var kNum int

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	path = make([]int, 0)
	kNum = k
	return path
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

}
