package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

var g [][]int
var sum [70][70]int

func construct(grid [][]int) *Node {
	g = grid
	return dfs(0, 0, len(g)-1, len(g)-1)
}

func dfs(a, b, c, d int) *Node {
	ok := true
	t := g[a][b]
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			if g[i][j] != t {
				ok = false
			}
		}
	}
	if ok {
		return &Node{Val: t == 1, IsLeaf: true}
	}
	root := &Node{Val: t == 1, IsLeaf: false}
	dx, dy := c-a+1, d-b+1
	root.TopLeft = dfs(a, b, a+dx/2-1, b+dy/2-1)
	root.TopRight = dfs(a, b+dy/2, a+dx/2-1, d)
	root.BottomLeft = dfs(a+dx/2, b, c, b+dy/2-1)
	root.BottomRight = dfs(a+dx/2, b+dy/2, c, d)
	return root
}
