package main

import "container/heap"

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

func minStoneSum(piles []int, k int) (ans int) {
	var hp IntHeap = IntHeap(piles)
	h := &hp
	heap.Init(h)
	for k > 0 {
		x := heap.Pop(h).(int)
		heap.Push(h, x-x/2)
		k--
	}
	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		ans += x
	}
	return ans
}
