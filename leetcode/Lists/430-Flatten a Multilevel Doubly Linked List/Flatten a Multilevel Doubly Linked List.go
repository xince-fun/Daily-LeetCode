package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten1(root *Node) *Node {
	dummyHead := &Node{Next: root}
	for root != nil {
		if root.Child != nil {
			tmp := root.Next
			chead := flatten1(root.Child)
			root.Next = chead
			chead.Prev = root
			root.Child = nil
			for root.Next != nil {
				root = root.Next
			}
			root.Next = tmp
			if tmp != nil {
				tmp.Prev = root
			}
			root = tmp
		} else {
			root = root.Next
		}
	}
	return dummyHead.Next
}

func flatten2(root *Node) *Node {
	return dfs(root)
}

func dfs(root *Node) *Node {
	last := root
	for root != nil {
		if root.Child == nil {
			last = root
			root = root.Next
		} else {
			tmp := root.Next
			childLast := dfs(root.Child)
			root.Next = root.Child
			root.Child.Prev = root
			root.Child = nil
			if childLast != nil {
				childLast.Next = tmp
			}
			if tmp != nil {
				tmp.Prev = childLast
			}
			last = root
			root = childLast
		}
	}
	return last
}

func flatten(root *Node) *Node {
	dummyHead := &Node{Next: root}

	for root != nil {
		if root.Child != nil {
			tmp := root.Next
			child := root.Child
			root.Next = child
			child.Prev = root
			root.Child = nil
			last := root
			for last.Next != nil {
				last = last.Next
			}
			last.Next = tmp
			if tmp != nil {
				tmp.Prev = last
			}
		}
		root = root.Next
	}
	return dummyHead.Next
}
