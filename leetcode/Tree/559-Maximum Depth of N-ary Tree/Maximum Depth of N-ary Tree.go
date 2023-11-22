package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxD := 0
	for _, child := range root.Children {
		temp := maxDepth(child)
		if temp > maxD {
			maxD = temp
		}
	}
	return maxD + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
