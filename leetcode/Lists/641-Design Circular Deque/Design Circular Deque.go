package main

type MyCircularDeque struct {
	Head, Tail int
	Cap        int
	Len        int
	Nums       []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{Cap: k, Nums: make([]int, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Head = (this.Head + this.Cap - 1) % this.Cap
	this.Nums[this.Head] = value
	this.Len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Nums[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Cap
	this.Len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % this.Cap
	this.Len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Tail = (this.Tail + this.Cap - 1) % this.Cap
	this.Len--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[this.Head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Nums[(this.Tail+this.Cap-1)%this.Cap]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Len == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Len == this.Cap
}
