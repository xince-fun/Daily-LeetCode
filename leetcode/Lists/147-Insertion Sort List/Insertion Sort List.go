package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: 0}
	cur, pre := head, newHead
	for cur != nil {
		node := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = newHead
		cur = node
	}
	return newHead.Next
}

func main() {
	n1 := &ListNode{Val: 4}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	fmt.Println(insertionSortList(n1))
}
