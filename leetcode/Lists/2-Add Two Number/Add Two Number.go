package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	n1, n2, carry := 0, 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

func main() {
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 5}
	n5 := &ListNode{Val: 6}
	n6 := &ListNode{Val: 4}

	n1.Next = n2
	n1.Next.Next = n3

	n4.Next = n5
	n4.Next.Next = n6

	fmt.Println(addTwoNumbers(n1, n4))
}
