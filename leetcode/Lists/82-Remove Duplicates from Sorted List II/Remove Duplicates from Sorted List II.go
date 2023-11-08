package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	p, q := dummyHead, head
	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			for q.Next != nil && q.Next.Val == q.Val {
				q = q.Next
			}
			q = q.Next
			if q == nil {
				p.Next = nil
			}
		} else {
			p.Next = q
			p = p.Next
			q = q.Next
		}
	}
	return dummyHead.Next
}

func main() {

}
