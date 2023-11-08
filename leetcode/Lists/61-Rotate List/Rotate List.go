package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	current := dummyHead
	length := 0
	for current.Next != nil {
		length++
		current = current.Next
	}
	if k%length == 0 {
		return head
	}
	current.Next = head
	current = dummyHead
	for i := length - k%length; i > 0; i-- {
		current = current.Next
	}
	res := &ListNode{Val: 0, Next: current.Next}
	current.Next = nil
	return res.Next
}

func main() {

}
