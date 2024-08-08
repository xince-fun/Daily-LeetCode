package main

import "fmt"

// 判断是否有相邻的 0
// 把 i 取反（保留低n位），计作 x, 判断 x 中是否有相邻的1
// x & (x >> 1)

func validStrings(n int) (ans []string) {
	for i := 0; i < 1<<n; i++ {
		x := 1<<n - 1 ^ i
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, i))
		}
	}
	return
}
