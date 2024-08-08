package main

import "container/heap"

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

const mod = 1e9 + 7

func maximumProduct(nums []int, k int) int {
	var hp IntHeap = IntHeap(nums)
	h := &hp
	heap.Init(h)
	for k > 0 {
		x := heap.Pop(h).(int)
		x += 1
		heap.Push(h, x)
		k--
	}
	ans := 1
	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		ans *= x
	}
	return ans % mod
}
