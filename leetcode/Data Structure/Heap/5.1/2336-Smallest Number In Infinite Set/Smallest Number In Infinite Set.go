package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type SmallestInfiniteSet struct {
	h   *IntHeap
	set map[int]bool
	idx int
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		h:   h,
		set: make(map[int]bool),
		idx: 1,
	}
}

func (t *SmallestInfiniteSet) PopSmallest() int {
	ans := -1
	if t.h.Len() > 0 {
		ans = heap.Pop(t.h).(int)
		t.set[ans] = false
	} else {
		ans = t.idx
		t.idx += 1
	}
	return ans
}

func (t *SmallestInfiniteSet) AddBack(num int) {
	if num >= t.idx || t.set[num] {
		return
	}
	if num == t.idx-1 {
		t.idx -= 1
	} else {
		heap.Push(t.h, num)
		t.set[num] = true
	}
}
