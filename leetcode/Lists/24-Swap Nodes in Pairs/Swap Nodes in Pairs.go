package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	current := dummyHead
	for current.Next != nil && current.Next.Next != nil {
		node1 := current.Next
		node2 := current.Next.Next
		current.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		current = node1
	}
	return dummyHead.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	for pt := dummyHead; pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummyHead.Next
}
