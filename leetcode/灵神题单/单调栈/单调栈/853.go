package main

import (
	"sort"
)

// Pos A < Pos B, A 先到达， 那AB形成车队
type zip struct {
	a []int
	b []int
}

func (p zip) Less(i, j int) bool { return p.a[i] < p.a[j] }
func (p zip) Len() int           { return len(p.a) }
func (p zip) Swap(i, j int) {
	p.a[i], p.a[j] = p.a[j], p.a[i]
	p.b[i], p.b[j] = p.b[j], p.b[i]
}

func carFleet(target int, position []int, speed []int) int {
	sort.Sort(zip{position, speed})
	n := len(position)
	time := make([]float64, n)
	for i := 0; i < n; i++ {
		time[i] = float64(target-position[i]) / float64(speed[i])
	}
	st := []float64{}
	for _, t := range time {
		for len(st) > 0 && t >= st[len(st)-1] {
			st = st[:len(st)-1]
		}
		st = append(st, float64(t))
	}
	return len(st)
}
