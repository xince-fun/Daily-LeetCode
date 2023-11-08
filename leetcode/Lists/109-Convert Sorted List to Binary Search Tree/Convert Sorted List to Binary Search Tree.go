package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode
type TreeNode = structure.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
	}
	mid, pre := midAndPreNode(head)
	if mid == nil {
		return nil
	}
	if pre != nil {
		pre.Next = nil
	}
	if mid == head {
		head = nil
	}
	return &TreeNode{Val: mid.Val, Left: sortedListToBST(head), Right: sortedListToBST(mid.Next)}
}

func midAndPreNode(head *ListNode) (mid *ListNode, pre *ListNode) {
	if head == nil || head.Next == nil {
		return nil, head
	}
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1, pre
}
