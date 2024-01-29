package main

func getDescentPeriods(prices []int) (ans int64) {
	for i, n := 0, len(prices); i < n; {
		i0 := i
		for i++; i < n && prices[i] == prices[i-1]-1; i++ {
		}
		ans += int64(i-i0) * int64(i-i0+1) / 2
	}
	return
}

func cal(a int) (ans int) {
	for i := a; i >= 1; i-- {
		ans *= i
	}
	return
}
