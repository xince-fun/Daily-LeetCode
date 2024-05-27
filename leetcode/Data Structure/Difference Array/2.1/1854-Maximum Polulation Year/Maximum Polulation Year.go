package main

func maximumPopulation(logs [][]int) (ans int) {
	d := [101]int{}
	for _, l := range logs {
		d[l[0]-1950]++
		d[l[1]-1950]--
	}
	s := 0
	m := 0
	for i, v := range d {
		s += v
		if s > m {
			m = s
			ans = i
		}
	}
	return ans + 1950
}
