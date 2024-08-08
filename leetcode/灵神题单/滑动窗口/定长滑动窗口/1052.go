package main

func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	sum := 0
	// 前minutes分钟不生气
	for i := 0; i < len(customers); i++ {
		if i < minutes {
			sum += customers[i]
		} else {
			sum += customers[i] * (grumpy[i] ^ 1)
		}
	}
	ans = sum
	for i := minutes; i < len(customers); i++ {
		sum += customers[i] * grumpy[i]
		sum -= customers[i-minutes] * grumpy[i-minutes]
		ans = max(ans, sum)
	}
	return ans
}
