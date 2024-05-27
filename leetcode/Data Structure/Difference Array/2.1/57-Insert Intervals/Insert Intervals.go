package main

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	d := make([]int, 200007)
	edge := 0

	for _, t := range intervals {
		d[t[0]*2]++
		d[t[1]*2+1]--
		edge = max(edge, t[1]*2)
	}
	d[newInterval[0]*2]++
	d[newInterval[1]*2+1]--
	edge = max(edge, newInterval[1]*2)

	s := 0
	start := -1
	for i := 0; i < edge+2; i++ {
		s += d[i]
		if s > 0 && start == -1 {
			start = i / 2
		} else if s == 0 && start != -1 {
			res = append(res, []int{start, i / 2})
			start = -1
		}
	}
	return
}

func insert1(intervals [][]int, newInterval []int) (res [][]int) {
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return
}

func insert2(intervals [][]int, newInterval []int) (res [][]int) {
	mn := 200002
	mx := 0

	for _, t := range intervals {
		mn = min(mn, t[0])
		mx = max(mx, t[1])
	}
	mn = min(newInterval[0], mn)
	mx = max(newInterval[1], mx)
	d := make([]int, (mx-mn+1)*2)

	for _, t := range intervals {
		d[(t[0]-mn)*2]++
		d[(t[1]-mn)*2+1]--
	}
	d[(newInterval[0]-mn)*2]++
	d[(newInterval[1]-mn)*2+1]--

	s := 0
	start := -1
	for i, v := range d {
		s += v
		if s > 0 && start == -1 {
			start = i/2 + mn
		} else if s == 0 && start != -1 {
			res = append(res, []int{start, i/2 + mn})
			start = -1
		}
	}
	return
}
