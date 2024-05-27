package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (t MyCalendarThree) add(x, delta int) {
	if v, ok := t.Get(x); ok {
		delta += v.(int)
	}
	t.Put(x, delta)
}

func (this *MyCalendarThree) Book(startTime int, endTime int) (ans int) {
	this.add(startTime, 1)
	this.add(endTime, -1)

	maxBook := 0
	for it := this.Iterator(); it.Next(); {
		maxBook += it.Value().(int)
		if maxBook > ans {
			ans = maxBook
		}
	}
	return
}
