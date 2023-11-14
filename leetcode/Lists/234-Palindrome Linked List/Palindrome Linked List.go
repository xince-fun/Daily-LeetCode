package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转
	preMiddle, preCurrent := slow, slow.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}

	slow, fast = head, preMiddle.Next
	res := true
	for slow != preMiddle {
		if slow.Val == fast.Val {
			slow = slow.Next
			fast = fast.Next
		} else {
			res = false
			break
		}
	}
	if slow == preMiddle {
		if fast != nil && slow.Val != fast.Val {
			return false
		}
	}
	return res
}

func isPalindrome1(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
