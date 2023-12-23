package main

import (
	"container/heap"
	"math"
)

type heapInt []int

func (h heapInt) Len() int {
	return len(h)
}

func (h heapInt) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h heapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minStoneSum(piles []int, k int) (ans int) {
	h := heapInt(piles)
	heap.Init(&h)

	for i := 0; i < k; i++ {
		x := heap.Pop(&h).(int)
		y := math.Floor(float64(x) / 2)
		heap.Push(&h, x-int(y))
	}

	for h.Len() > 0 {
		x := heap.Pop(&h).(int)
		ans += x
	}
	return
}
