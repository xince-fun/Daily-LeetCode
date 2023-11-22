package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

var maxNum int
var m map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	ans := make([]int, 0)
	maxNum = 0
	m = make(map[int]int)
	travel(root)
	for k, v := range m {
		if v == maxNum {
			ans = append(ans, k)
		}
	}
	return ans
}

func travel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := root.Val + travel(root.Left) + travel(root.Right)
	_, ok := m[root.Val]
	if ok {
		m[cur] = m[cur] + 1
	} else {
		m[cur] = 1
	}
	maxNum = max(maxNum, m[cur])
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
