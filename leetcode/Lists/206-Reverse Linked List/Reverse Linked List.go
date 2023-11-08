package main

import (
	"fmt"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

// 朴素双指针
func reverseList1(head *ListNode) *ListNode {
	current := head
	tmp := &ListNode{}
	var pre *ListNode
	for current != nil {
		tmp = current.Next // 下一个
		current.Next = pre
		pre = current
		current = tmp
	}
	return pre
}

// 精简版
func reverseList2(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

// 递归1
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp)
}

func reverseList3(head *ListNode) *ListNode {
	return reverse(nil, head)
}

// 递归2
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println(reverseList(n1))
}
