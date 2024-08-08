package main

func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)
	left, right := -1, n+1
	check := func(mid int) bool {
		del := make([]bool, len(s))
		for _, x := range removable[:mid] {
			del[x] = true
		}
		i, j := 0, 0
		for i < len(s) && j < len(p) {
			if !del[i] && s[i] == p[j] {
				j++
			}
			i++
		}
		return j == len(p)
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
