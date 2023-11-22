package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return quadTree1
	}
	tl, tr, bl, br := intersect(quadTree1.TopLeft, quadTree2.TopLeft), intersect(quadTree1.TopRight, quadTree2.TopRight),
		intersect(quadTree1.BottomLeft, quadTree2.BottomLeft), intersect(quadTree1.BottomRight, quadTree2.BottomRight)

	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tl.Val == bl.Val && tl.Val == br.Val {
		return &Node{Val: tl.Val, IsLeaf: true}
	}
	return &Node{false, false, tl, tr, bl, br}
}
