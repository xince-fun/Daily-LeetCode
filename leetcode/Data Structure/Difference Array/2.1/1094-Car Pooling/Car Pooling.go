package main

import "container/heap"

type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func carPooling1(trips [][]int, capacity int) bool {
	h := IntHeap{}

	heap.Init(&h)
	count := 0
	for i := 0; i < len(trips); i++ {
		trip := trips[i]

		for len(h) > 0 && h[0][2] <= trip[1] {
			count -= (heap.Pop(&h).([]int))[0]
		}
		count += trip[0]
		if count > capacity {
			return false
		}
		heap.Push(&h, trip)
	}
	return true
}

func carPooling(trips [][]int, capacity int) bool {
	d := [1001]int{}
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0]
	}
	s := 0
	for _, v := range d {
		s += v
		if s > capacity {
			return false
		}
	}
	return true
}
