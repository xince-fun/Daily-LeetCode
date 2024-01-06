package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		val := gcd(cur.Val, cur.Next.Val)
		cur.Next = &ListNode{Val: val, Next: tmp}
		cur = tmp
	}
	return head
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
