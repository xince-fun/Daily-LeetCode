package main

import (
	"fmt"
	"strconv"

	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var h, m, n int
var ans [][]string

func printTree(root *TreeNode) [][]string {
	h, m, n = 0, 0, 0
	maxDepth(root, 0)
	m, n = h+1, (1<<(h+1))-1
	ans = make([][]string, 0)
	for i := 0; i < m; i++ {
		temp := make([]string, n)

		ans = append(ans, temp)
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

func maxDepth(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	h = max(h, depth)
	maxDepth(root.Left, depth+1)
	maxDepth(root.Right, depth+1)
}

func dfs(root *TreeNode, x, y int) {
	if root == nil {
		return
	}
	ans[x][y] = strconv.Itoa(root.Val)
	if h-x-1 >= 0 {
		dfs(root.Left, x+1, y-(1<<(h-x-1)))
	}
	if h-x-1 >= 0 {
		dfs(root.Right, x+1, y+(1<<(h-x-1)))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 1}
	root.Left = left
	fmt.Println(printTree(root))
}
