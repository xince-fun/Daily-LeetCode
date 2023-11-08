package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			index1, index2 := fast, head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index2
		}
	}
	return nil
}
