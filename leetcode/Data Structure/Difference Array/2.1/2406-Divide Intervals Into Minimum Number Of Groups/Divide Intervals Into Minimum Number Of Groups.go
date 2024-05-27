package main

import "github.com/emirpasic/gods/trees/redblacktree"

func minGroups(intervals [][]int) (ans int) {
	d := redblacktree.NewWithIntComparator()

	for _, t := range intervals {
		add(d, t[0], 1)
		add(d, t[1]+1, -1)
	}
	s := 0

	for it := d.Iterator(); it.Next(); {
		s += it.Value().(int)
		ans = max(ans, s)
	}
	return
}

func add(d *redblacktree.Tree, x, delta int) {
	if v, ok := d.Get(x); ok {
		delta += v.(int)
	}
	d.Put(x, delta)
}

func minGroups1(intervals [][]int) (ans int) {
	mn := 1000001
	mx := 0
	for _, t := range intervals {
		mn = min(mn, t[0])
		mx = max(mx, t[1])
	}
	diff := make([]int, mx-mn+1)
	for _, t := range intervals {
		diff[t[0]-mn]++
		if t[1] < mx {
			diff[t[1]-mn+1]--
		}
	}
	s := 0
	for _, v := range diff {
		s += v
		ans = max(ans, s)
	}
	return
}
