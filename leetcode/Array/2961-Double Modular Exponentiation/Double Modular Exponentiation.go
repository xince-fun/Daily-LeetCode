package main

// ((a^b % 10)^c) % m

// ((2^4 % 10)^3) % 10
// =

// ((3^3 % 3)^3) % 2

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func getGoodIndices(variables [][]int, target int) []int {
	ans := []int{}
	for i, v := range variables {
		a, b, c, m := v[0], v[1], v[2], v[3]
		if pow(pow(a, b, 10), c, m) == target {
			ans = append(ans, i)
		}
	}
	return ans
}
