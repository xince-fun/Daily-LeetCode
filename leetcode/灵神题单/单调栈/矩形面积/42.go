package main

// 找上一个最大元素，在找的过程中填坑

func trap(height []int) (ans int) {
	st := []int{}
	for i, h := range height {
		for len(st) > 0 && h >= height[st[len(st)-1]] {
			bottomH := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			dh := min(height[left], h) - bottomH
			ans += dh * (i - left - 1)
		}
		st = append(st, i)
	}
	return
}
