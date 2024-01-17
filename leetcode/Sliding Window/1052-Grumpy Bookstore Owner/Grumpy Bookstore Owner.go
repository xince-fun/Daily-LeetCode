package main

func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			ans += customers[i]
			customers[i] = 0
		}
	}
	sum, cur := 0, 0
	for i := 0; i < n; i++ {
		cur += customers[i]
		if i >= minutes {
			cur -= customers[i-minutes]
		}
		sum = max(sum, cur)
	}
	return ans + sum
}
