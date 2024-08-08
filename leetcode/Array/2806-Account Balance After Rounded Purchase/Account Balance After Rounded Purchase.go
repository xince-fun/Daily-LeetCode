package main

func accountBalanceAfterPurchase(purchaseAmount int) int {
	m := purchaseAmount / 10
	n := purchaseAmount % 10
	if n >= 5 {
		return 100 - (m+1)*10
	}
	return 100 - m*10
}
