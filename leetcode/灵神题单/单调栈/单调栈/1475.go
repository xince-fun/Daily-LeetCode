package main

func finalPrices1(prices []int) []int {
	n := len(prices)
	st := []int{}
	ans := make([]int, n)
	for i, p := range prices {
		for len(st) > 0 && p <= prices[st[len(st)-1]] {
			x := st[len(st)-1]
			st = st[:len(st)-1]
			ans[x] = prices[x] - p
		}
		st = append(st, i)
	}
	for _, v := range st {
		ans[v] = prices[v]
	}
	return ans
}

func finalPrices(prices []int) []int {
	n := len(prices)
	st := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 0 && p < prices[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = p - prices[st[len(st)-1]]
		} else {
			ans[i] = p
		}
		st = append(st, i)
	}
	return ans
}
