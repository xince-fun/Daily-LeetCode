package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func rob1(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfs(root.Left)
	r0, r1 := dfs(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 选 = 左不选 + 右不选 + 当前节点值
// 不选 = max(左选， 左不选) + max(右选， 右不选)

// 没有上司的舞会
// 一般树
// 选 = sum(不选子节点) + 当前节点
// 不选 = sum(选子节点，不选子节点)

// 最大独立集：需要从图中选出尽量多的点，使得这些点互不相邻。
// 常见套路： 选/不选、枚举选哪个

func rob(root *TreeNode) int {

	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		lRob, lNotRob := dfs(root.Left)
		rRob, rNotRob := dfs(root.Right)
		rob := lNotRob + rNotRob + root.Val
		notRob := max(lRob, lNotRob) + max(rRob, rNotRob)
		return rob, notRob
	}

	return max(dfs(root))
}
