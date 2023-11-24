package main

import (
	"github.com/xince-fun/daily-leetcode/structure"
)

type TreeNode = structure.TreeNode

func findTarget1(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	var check func(*TreeNode, int) bool
	check = func(root *TreeNode, k int) bool {
		if root == nil {
			return false
		}
		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		return check(root.Left, k) || check(root.Right, k)
	}
	return check(root, k)
}

func findTarget(root *TreeNode, k int) bool {
	lq, rq := []*TreeNode{}, []*TreeNode{}
	temp := root
	for temp != nil {
		lq = append(lq, temp)
		temp = temp.Left
	}
	temp = root
	for temp != nil {
		rq = append(rq, temp)
		temp = temp.Right
	}
	left, right := lq[len(lq)-1], rq[len(rq)-1]
	for left.Val < right.Val {
		t := left.Val + right.Val
		if t == k {
			return true
		}
		if t < k {
			left = getNext(&lq, true)
		} else {
			right = getNext(&rq, false)
		}
	}
	return false
}

func getNext(q *[]*TreeNode, isLeft bool) *TreeNode {
	var node *TreeNode
	if isLeft {
		node = (*q)[len(*q)-1].Right
	} else {
		node = (*q)[len(*q)-1].Left
	}
	*q = (*q)[:len(*q)-1]
	for node != nil {
		*q = append(*q, node)
		if isLeft {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return (*q)[len(*q)-1]
}
