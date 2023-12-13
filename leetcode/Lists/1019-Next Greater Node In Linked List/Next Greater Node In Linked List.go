package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func nextLargerNodes1(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(stack) != 0 && x >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = nums[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return ans
}

func nextLargerNodes2(head *ListNode) (ans []int) {
	st := []int{} // 单调栈
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
	f(head, 0)
	return
}

// 底大顶小 从左往右看
func nextLargerNodes(head *ListNode) (ans []int) {
	type pair struct{ x, i int }
	st := []pair{} // (节点值，节点下标)
	for cur := head; cur != nil; cur = cur.Next {
		x := cur.Val
		for len(st) > 0 && st[len(st)-1].x < x {
			ans[st[len(st)-1].i] = x
			st = st[:len(st)-1]
		}
		st = append(st, pair{x, len(ans)})
		ans = append(ans, 0)
	}
	return ans
}
