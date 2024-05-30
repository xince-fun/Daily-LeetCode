package main

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	diff := map[int]int{}
	for _, f := range flowers {
		diff[f[0]]++
		diff[f[1]+1]--
	}

	n := len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	id := make([]int, len(people))
	for i := range id {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool { return people[id[i]] < people[id[j]] })

	j, sum := 0, 0

	for _, i := range id {
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]]
		}
		people[i] = sum
	}
	return people
}

func add(d *redblacktree.Tree, x, delta int) {
	if v, ok := d.Get(x); ok {
		delta += v.(int)
	}
	d.Put(x, delta)
}
