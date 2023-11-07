package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	for list1 != nil && list2 != nil {
		n := 0
		if list1.Val < list2.Val {
			n = list1.Val
			list1 = list1.Next
		} else {
			n = list2.Val
			list2 = list2.Next
		}
		current.Next = &ListNode{Val: n}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return dummyHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
