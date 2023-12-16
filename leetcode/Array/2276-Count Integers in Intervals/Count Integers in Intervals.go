package main

import "container/heap"

type sliceHeap [][]int

func (h sliceHeap) Len() int {
	return len(h)
}

func (h sliceHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h sliceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *sliceHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *sliceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type CountIntervals struct {
	h *sliceHeap
}

func Constructor() CountIntervals {
	h := &sliceHeap{}
	heap.Init(h)
	return CountIntervals{h: h}
}

func (this *CountIntervals) Add(left int, right int) {
	arr := []int{left, right}
	this.h.Push(arr)
	res := [][]int{}
	interval := this.h.Pop().([]int)
	res = append(res, interval)
	for this.h.Len() > 0 {
		interval := this.h.Pop().([]int)
		if res[len(res)-1][1] >= interval[0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		} else {
			res = append(res, interval)
		}
	}
	for _, r := range res {
		this.h.Push(r)
	}
}

func (this *CountIntervals) Count() int {
	ans := 0
	temp := [][]int{}
	for this.h.Len() > 0 {
		interval := this.h.Pop().([]int)
		ans += interval[1] - interval[0] + 1
		temp = append(temp, interval)
	}
	for _, t := range temp {
		this.h.Push(t)
	}
	return ans
}
