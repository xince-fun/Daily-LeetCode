package main

import "sort"

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.m[key] = append(t.m[key], pair{timestamp, value})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	pairs := t.m[key]
	idx := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if idx > 0 {
		return pairs[idx-1].value
	}
	return ""
}
