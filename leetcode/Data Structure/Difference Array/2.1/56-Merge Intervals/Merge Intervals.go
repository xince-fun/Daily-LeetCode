package main

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func merge(intervals [][]int) (res [][]int) {
	d := make([]int, 20007)
	edge := 0

	for _, t := range intervals {
		d[t[0]*2]++
		d[t[1]*2+1]--
		edge = max(edge, t[1]*2)
	}

	start := -1
	s := 0
	for i := 0; i < edge; i++ {
		s += d[i]
		if s > 0 && start == -1 {
			start = i / 2
		}
		if s == 0 && start != -1 {
			res = append(res, []int{start, i / 2})
			start = -1
		}
	}
	return
}

func merge1(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else if intervals[i][0] <= res[len(res)-1][1] {
			if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		}
	}
	return
}

func merge2(intervals [][]int) (res [][]int) {
	d := redblacktree.NewWithIntComparator()

	for _, t := range intervals {
		delta := 1
		if v, ok := d.Get(t[0]); ok {
			delta += v.(int)
		}
		d.Put(t[0], delta)
		delta = -1
		if v, ok := d.Get(t[1]); ok {
			delta += v.(int)
		}
		d.Put(t[1], delta)
	}

	s := 0
	left := 0
	for it := d.Iterator(); it.Next(); {
		if s == 0 {
			left = it.Key().(int)
		}
		s += it.Value().(int)
		if s == 0 {
			res = append(res, []int{left, it.Key().(int)})
		}
	}
	return
}

func merge3(intervals [][]int) (res [][]int) {
	mn := 200002
	mx := 0

	for _, t := range intervals {
		mn = min(mn, t[0])
		mx = max(mx, t[1])
	}

	d := make([]int, (mx-mn+1)*2)

	for _, t := range intervals {
		d[(t[0]-mn)*2]++
		d[(t[1]-mn)*2+1]--
	}

	start := -1
	s := 0

	for i, v := range d {
		s += v
		if s > 0 && start == -1 {
			start = i/2 + mn
		}
		if s == 0 && start != -1 {
			res = append(res, []int{start, i/2 + mn})
			start = -1
		}
	}

	return
}
