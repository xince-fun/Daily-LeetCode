package main

import "sort"

type pair struct {
	spanId int
	value  int
}

type SnapshotArray struct {
	curSpanId int
	history   map[int][]pair
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		history: make(map[int][]pair),
	}
}

func (t *SnapshotArray) Set(index int, val int) {
	t.history[index] = append(t.history[index], pair{t.curSpanId, val})
}

func (t *SnapshotArray) Snap() int {
	t.curSpanId++
	return t.curSpanId - 1
}

func (t *SnapshotArray) Get(index int, snap_id int) int {
	h := t.history[index]
	// 找快照编号<=spanId的最后一次修改
	// 等于找 >= spanId +1, 它的上一个
	j := sort.Search(len(h), func(j int) bool { return h[j].spanId >= snap_id+1 }) - 1
	if j >= 0 {
		return h[j].value
	}
	return 0
}
