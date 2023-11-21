package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*Node{root}
	for len(q) > 0 {
		length := len(q)
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			for _, n := range node.Children {
				if n != nil {
					q = append(q, n)
				}
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
