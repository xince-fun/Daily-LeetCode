package main

import (
	"container/heap"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

// 解法1: 暴力解法
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return &ListNode{}
	}
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	end := false
	for end != true {
		index, n := 0, 10001
		find := false
		for i, l := range lists {
			if l != nil && l.Val < n {
				n = l.Val
				index = i
				find = true
			}
		}
		if !find {
			break
		}
		lists[index] = lists[index].Next
		current.Next = &ListNode{Val: n}
		current = current.Next
	}
	return dummyHead.Next
}

// 解法2:最小堆
type hp []*ListNode

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists2(lists []*ListNode) *ListNode {
	h := hp{}

	for _, l := range lists {
		if l != nil {
			h = append(h, l)
		}
	}
	heap.Init(&h)

	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		current.Next = node
		current = current.Next
	}
	return dummyHead.Next
}

// 解法3: 分治
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:m/2])
	right := mergeKLists(lists[m/2:])
	return mergeTwoLists(left, right)
}
