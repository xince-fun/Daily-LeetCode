package main

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	var gh IntHeap = IntHeap(gifts)
	h := &gh
	heap.Init(h)
	for k > 0 {
		g := heap.Pop(h)
		x := g.(int)
		left := math.Sqrt(float64(x))
		heap.Push(h, int(left))
		k--
	}
	sum := 0
	for h.Len() > 0 {
		sum += heap.Pop(h).(int)
	}
	return int64(sum)
}
