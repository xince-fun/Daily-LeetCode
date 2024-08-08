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

type SeatManager struct {
	h   *IntHeap
	idx int
	set map[int]bool
}

func Constructor(n int) SeatManager {
	h := &IntHeap{}
	heap.Init(h)
	return SeatManager{
		h:   h,
		idx: 1,
		set: map[int]bool{},
	}
}

func (t *SeatManager) Reserve() int {
	ans := -1
	if t.h.Len() > 0 {
		ans = heap.Pop(t.h).(int)
	} else {
		ans = t.idx
		t.idx += 1
	}
	t.set[ans] = true // 不可预约
	return ans
}

func (t *SeatManager) Unreserve(seatNumber int) {
	if seatNumber >= t.idx || !t.set[seatNumber] {
		return
	}
	if seatNumber == t.idx-1 {
		t.idx -= 1
		t.set[t.idx] = false
	} else {
		heap.Push(t.h, seatNumber)
		t.set[seatNumber] = false
	}
}
