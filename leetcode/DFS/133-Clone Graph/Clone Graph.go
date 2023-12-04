package main

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
	visited = make(map[*Node]*Node)
	if node == nil {
		return nil
	}
	return dfs(node)
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(n))
	}
	return newNode
}
