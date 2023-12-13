package main

func finalPrices1(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && prices[stack[len(stack)-1]] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = prices[i]
		} else {
			ans[i] = prices[i] - prices[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return ans
}

func finalPrices(prices []int) []int {
	n := len(prices)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && prices[stack[len(stack)-1]] < prices[i] {
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prices[i] -= prices[x]
		}
		stack = append(stack, i)
	}
	return prices
}
