package main

func maxProfit(prices []int) (ans int) {
	// 从左往右枚举
	minPrices := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrices)
		minPrices = min(minPrices, p)
	}
	return
}
