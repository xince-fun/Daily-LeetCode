package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummyHead, dummyBig := &ListNode{}, &ListNode{}
	pre, cur, q := dummyHead, head, dummyBig
	for cur != nil {
		if cur.Val < x {
			pre.Next = cur
			pre = pre.Next
			cur = cur.Next
		} else {
			q.Next = cur
			q = q.Next
			cur = cur.Next
		}
	}
	q.Next = nil
	pre.Next = dummyBig.Next
	return dummyHead.Next
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	x := 3
	fmt.Println(partition(n1, x))
}
