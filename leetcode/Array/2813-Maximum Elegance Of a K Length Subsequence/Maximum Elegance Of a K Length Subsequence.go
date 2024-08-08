package main

import "slices"

/*
	1 <= items.length == n <= 105
	items[i].length == 2
	items[i][0] == profiti
	items[i][1] == categoryi
	1 <= profiti <= 109
	1 <= categoryi <= n
	1 <= k <= n
*/

func findMaximumElegance(items [][]int, k int) int64 {
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := []int{}
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if !vis[category] {
				vis[category] = true
			} else {
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] {
			vis[category] = true //
			totalProfit += profit - duplicate[len(duplicate)-1]
			duplicate = duplicate[:len(duplicate)-1]
		}
		ans = max(ans, totalProfit+len(vis)*len(vis))

	}
	return int64(ans)
}
