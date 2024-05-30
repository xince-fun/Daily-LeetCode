package main

func splitPainting(segments [][]int) (res [][]int64) {
	d := make([]int, 100002)

	for _, s := range segments {
		d[s[0]] += s[2]
		d[s[1]] -= s[2]
	}

	sum := 0
	start := -1
	for i, v := range d {
		now := sum
		if sum > 0 && start == -1 {
			start = i
		}
		sum += v
		if now != sum && now > 0 {
			res = append(res, []int64{int64(start), int64(i), int64(now)})
		}
	}
	return
}
