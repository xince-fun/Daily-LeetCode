package main

import (
	"container/heap"
	"sort"
)

/*
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

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	type pair struct{ x, y int }
	arrive := []pair{}
	leave := []pair{}
	for i, time := range times {
		arrive = append(arrive, pair{time[0], i})
		leave = append(leave, pair{time[1], i})
	}
	sort.Slice(arrive, func(i, j int) bool {
		return arrive[i].x < arrive[j].x
	})
	sort.Slice(leave, func(i, j int) bool {
		return leave[i].x < leave[j].x
	})

	h := &IntHeap{}
	for i := 0; i < n; i++ {
		*h = append(*h, i)
	}
	heap.Init(h)
	set := map[int]int{}
	idx := 0

	for _, p := range arrive {
		// 到达时间与之前的离开操作
		for idx < n && leave[idx].x <= p.x {
			heap.Push(h, set[leave[idx].y])
			idx++
		}
		// 到达
		if h.Len() > 0 {
			x := heap.Pop(h).(int)
			set[p.y] = x
			if p.y == targetFriend {
				return x
			}
		}
	}
	return -1
}

*/

func smallestChair(times [][]int, targetFriend int) int {
	events := make([][2][]int, 1e5+1)

	for i, t := range times {
		l, r := t[0], t[1]
		events[l][1] = append(events[l][1], i) // 朋友到达
		events[r][0] = append(events[r][0], i) // 朋友离开
	}

	n := len(times)
	unoccupied := hp{make([]int, n)}
	for i := range unoccupied.IntSlice {
		unoccupied.IntSlice[i] = i
	}

	belong := make([]int, n)
	for _, e := range events {
		for _, id := range e[0] { // 处理离开的
			heap.Push(&unoccupied, belong[id])
		}
		for _, id := range e[1] {
			belong[id] = heap.Pop(&unoccupied).(int)
			if id == targetFriend {
				return belong[id]
			}
		}
	}

	return 0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = h.IntSlice[:len(a)-1]
	return v
}
