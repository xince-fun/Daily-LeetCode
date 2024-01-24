package main

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	suf := make([]int, n+1)
	st := []int{n} // 一个哨兵节点
	sum := 0
	for i := n - 1; i >= 0; i-- {
		x := maxHeights[i]
		for len(st) > 1 && x <= maxHeights[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			sum -= maxHeights[j] * (st[len(st)-1] - j)
		}
		sum += x * (st[len(st)-1] - i)
		suf[i] = sum
		st = append(st, i)
	}
	ans := sum

	st = []int{-1}
	pre := 0
	for i, x := range maxHeights {
		for len(st) > 1 && x <= maxHeights[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			pre -= maxHeights[j] * (j - st[len(st)-1])
		}
		pre += x * (i - st[len(st)-1])
		ans = max(ans, pre+suf[i+1])
		st = append(st, i)
	}
	return int64(ans)
}
