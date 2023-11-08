package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode
type TreeNode = structure.TreeNode

func sortedListToBST1(head *ListNode) *TreeNode {
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
	return &TreeNode{Val: mid.Val, Left: sortedListToBST1(head), Right: sortedListToBST1(mid.Next)}
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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	if pre != nil {
		pre.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}
