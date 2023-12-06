package main

import (
	"fmt"
)

func countGoodStrings(low int, high int, zero int, one int) int {
	ans := 0
	// 回溯
	const MOD = 1_000_000_007
	// 算每个长度为 i 的字符串的个数
	// 就是爬楼梯 f[i] 表示 i 的字符串的个数
	// f[i] = f[i-zero] + f[i-one]
	f := make([]int, high+1)
	f[0] = 1 // 空串的方案数为1
	for i := 1; i <= high; i++ {
		if i >= one {
			f[i] = (f[i] + f[i-one]) % MOD
		}
		if i >= zero {
			f[i] = (f[i] + f[i-zero]) % MOD
		}
		if i >= low {
			ans = (ans + f[i]) % MOD
		}
	}

	return ans
}

func main() {
	low, high := 200, 200
	zero, one := 10, 1
	fmt.Println(countGoodStrings(low, high, zero, one))
}
