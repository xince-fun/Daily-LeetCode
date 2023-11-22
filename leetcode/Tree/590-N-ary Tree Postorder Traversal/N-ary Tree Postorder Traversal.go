package main

type Node struct {
	Val      int
	Children []*Node
}

var ans []int

func postorder(root *Node) []int {
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
	for _, child := range root.Children {
		dfs(child)
	}
	ans = append(ans, root.Val)

}
