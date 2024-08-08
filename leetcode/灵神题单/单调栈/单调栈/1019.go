package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func nextLargerNodes1(head *ListNode) []int {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	ans := make([]int, len(arr))
	st := []int{}
	for i, x := range arr {
		for len(st) > 0 && x > arr[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			ans[idx] = x
		}
		st = append(st, i)
	}
	return ans
}

// 对于链表，从头开始递归，「归」的过程，相当于从右往左遍历
func nextLargerNodes(head *ListNode) (ans []int) {
	st := []int{}
	var f func(*ListNode, int)
	f = func(node *ListNode, i int) {
		if node == nil {
			ans = make([]int, i)
			return
		}
		f(node.Next, i+1)
		for len(st) > 0 && node.Val >= st[len(st)-1] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1]
		}
		st = append(st, node.Val)
	}
	return
}
