package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type Node = structure.Node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		len := len(q)
		for i := 0; i < len; i++ {
			cur := q[0]
			q = q[1:]
			if i+1 < len {
				cur.Next = q[0]
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}
