package main

/*
1. lowbit(n) = n & -n
2.
*/

const mod int = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit // n -= lowbit
	}
	// 1. 暴力
	// for _, q := range queries {
	// 	mul := 1
	// 	for _, x := range powers[q[0] : q[1]+1] {
	// 		mul = mul * x % mod
	// 	}
	// 	ans = append(ans, mul)
	// }
	np := len(powers)
	res := make([][]int, np)
	for i, x := range powers {
		res[i] = make([]int, np)
		res[i][i] = x
		for j := i + 1; j < np; j++ {
			res[i][j] = res[i][j-1] * powers[j] % mod
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}
