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

func minOperations(nums []int, k int) (ans int) {
	var hp IntHeap = IntHeap(nums)
	h := &hp
	heap.Init(h)
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x >= k && y >= k {
			break
		}
		heap.Push(h, min(x, y)*2+max(x, y))
		ans++
	}
	return
}
