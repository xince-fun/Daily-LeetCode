package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

// 单链表
func reorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//                 preMiddle
	//                     |
	// 反转链表后半部分 1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle, preCurrent := slow, slow.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	// 重新拼接 1->2-3->4->6->5->4 to 1->6->2->5->3->4
	slow, fast = head, preMiddle.Next
	for slow != preMiddle {
		preMiddle.Next = fast.Next
		fast.Next = slow.Next
		slow.Next = fast
		slow = fast.Next
		fast = preMiddle.Next
	}
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	arr := []*ListNode{}
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	length := len(arr)
	last := head
	cur = head
	for i := 0; i < length/2; i++ {
		tmp := &ListNode{Val: arr[length-i-1].Val, Next: cur.Next}
		cur.Next = tmp
		cur = tmp.Next
		last = tmp
	}
	if length%2 == 0 {
		last.Next = nil
	} else {
		cur.Next = nil
	}
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// cur 就是 head2
	for cur.Next != nil {
		next1 := head.Next
		next2 := cur.Next
		head.Next = cur
		cur.Next = next1
		head = next1
		cur = next2
	}
}

func main() {

}
