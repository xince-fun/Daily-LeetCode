package main

type Node struct {
	Val      int
	Children []*Node
}

var ans []int

func preorder(root *Node) []int {
	ans = make([]int, 0)
	if root == nil {
		return ans
	}
	dfs(root)
	return ans
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	for _, child := range root.Children {
		dfs(child)
	}

}
