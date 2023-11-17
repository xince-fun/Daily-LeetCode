package main

/*
type MyCircularQueue struct {
	Head, Tail *node
	Cap        int
	Len        int
}

type node struct {
	value int
	next  *node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Cap: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &node{value: value}
	if this.Head == nil {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.next = node
		this.Tail = node
	}
	this.Len++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = this.Head.next
	this.Len--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.value
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.value
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Len == this.Cap
}
*/

type MyCircularQueue struct {
	Head, Tail int
	Cap        int
	Nums       []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head: 0,
		Tail: 0,
		Cap:  k,
		Nums: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Nums[this.Tail%this.Cap] = value
	this.Tail = this.Tail + 1
	return this.Tail >= 0
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = this.Head + 1
	return this.Head >= 0
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[this.Head%this.Cap]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[(this.Tail-1)%this.Cap]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Tail-this.Head == this.Cap
}
