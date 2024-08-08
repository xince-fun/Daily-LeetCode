package main

func nthMagicalNumber(n int, a int, b int) int {
	const mod = 1e9 + 7
	lcm := a / gcd(a, b) * b
	// 第n个神奇数字是x
	// <=x的神奇数字有n个
	// x越大，n越大

	left, right := 0, min(a*b)*n
	for left+1 < right {
		mid := left + (right-left)/2
		if mid/a+mid/b-mid/lcm >= n {
			right = mid
		} else {
			left = mid
		}
	}
	return right % mod
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
