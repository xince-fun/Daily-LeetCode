package main

import "container/heap"

type FloatHeap []float64

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h FloatHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *FloatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}
func (h *FloatHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// sum - s >= sum / 2
// sum / 2 >= s

func halveArray(nums []int) (ans int) {
	sum := float64(0)
	h := &FloatHeap{}
	for _, x := range nums {
		sum += float64(x)
		*h = append(*h, float64(x))
	}
	heap.Init(h)
	s := float64(0)
	for h.Len() > 0 {
		if sum/2.0 <= s {
			break
		}
		x := heap.Pop(h).(float64)
		x = x / 2.0
		s += x
		heap.Push(h, x)
		ans++
	}
	return
}
