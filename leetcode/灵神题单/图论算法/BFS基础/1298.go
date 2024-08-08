package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) (ans int) {
	n := len(status)
	q := []int{}
	m := make([]bool, n)
	vis := make([]bool, n)

	for _, x := range initialBoxes {
		q = append(q, x)
		m[x] = true
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		// 没拿走 box 的物品，且可以打开
		if !vis[e] && status[e] == 1 {
			ans += candies[e]
			vis[e] = true
			q = append(q, containedBoxes[e]...)
			for _, key := range keys[e] {
				status[key] = 1
				if !vis[key] && m[key] {
					q = append(q, key)
				}
			}
		} else {
			m[e] = true
		}
	}
	return ans
}
