package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	for _, s := range stones {
		*h = append(*h, s)
	}
	heap.Init(h)
	for h.Len() > 1 {
		big := heap.Pop(h)
		lBig := heap.Pop(h)
		x := big.(int)
		y := lBig.(int)
		if x > y {
			heap.Push(h, x-y)
		} else if x == y {
			continue
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
