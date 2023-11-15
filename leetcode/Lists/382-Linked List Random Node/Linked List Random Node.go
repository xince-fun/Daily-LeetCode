package main

import (
	"math/rand"
	"time"

	"github.com/xince-fun/daily-leetcode/structure"
)

type ListNode = structure.ListNode

type Solution struct {
	head   *ListNode
	Length int
}

func Constructor(head *ListNode) Solution {
	cur := head
	length := 0
	for cur != nil {
		cur = cur.Next
		length++
	}
	return Solution{
		head:   head,
		Length: length,
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (this *Solution) GetRandom() int {
	randomNum := r.Intn(this.Length)
	cur := this.head
	for randomNum > 0 {
		cur = cur.Next
		randomNum--
	}
	return cur.Val
}
