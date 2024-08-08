package main

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: []int{},
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.q = append(r.q, t)
	for len(r.q) > 0 {
		if t-r.q[0] > 3000 {
			r.q = r.q[1:]
		} else {
			break
		}
	}
	return len(r.q)
}
