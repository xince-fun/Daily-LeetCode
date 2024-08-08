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

func maxKelements(nums []int, k int) (ans int64) {
	var mh IntHeap = IntHeap(nums)
	h := &mh
	heap.Init(h)
	for k > 0 {
		x := heap.Pop(h).(int)
		ans += int64(x)
		heap.Push(h, int(math.Ceil(float64(x)/3.0)))
		k--
	}
	return
}
