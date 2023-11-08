package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || left == right {
		return head
	}
	// 区间 [a.next, b] 反转
	dummyHead := &ListNode{Next: head}
	a, b := dummyHead, head
	// a 前面那个
	for i := 1; i < left; i++ {
		a = a.Next
	}
	for i := 1; i < right+1; i++ {
		b = b.Next
	}
	newHead := reverse(a.Next, b)
	a.Next = newHead
	return dummyHead.Next
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil || left == right {
		return head
	}
	newHead := &ListNode{Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

func reverse(a, b *ListNode) *ListNode {
	pre := b
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {}
