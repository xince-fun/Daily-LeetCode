package main

func isCovered(ranges [][]int, left int, right int) bool {
	d := [52]int{}
	for _, r := range ranges {
		d[r[0]]++
		d[r[1]+1]--
	}
	s := 0
	for _, v := range d[:left] {
		s += v
	}
	for _, v := range d[left : right+1] {
		if s += v; s == 0 {
			return false
		}
	}
	return true
}
