package main

func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left := make([]int, n)
	st := []int{}
	for i, x := range heights {
		for len(st) > 0 && x <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n)
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && heights[i] <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}
	for i, h := range heights {
		ans = max(ans, h*(right[i]-left[i]-1))
	}
	return ans
}
