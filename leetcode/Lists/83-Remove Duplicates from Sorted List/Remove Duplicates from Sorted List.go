package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
