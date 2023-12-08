package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k个待反转元素
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummyHead := &ListNode{Next: head}
	p0 := dummyHead

	for n >= k {
		n -= k
		var pre *ListNode
		cur = p0.Next
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		next := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = next
	}
	return dummyHead.Next
}
