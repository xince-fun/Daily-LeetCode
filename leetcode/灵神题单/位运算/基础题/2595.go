package main

import "math/bits"

func evenOddBit1(n int) []int {
	ans := make([]int, 2)
	for i := 0; i < bits.Len(uint(n)); i++ {
		ans[i&1] += n >> i & 1
	}
	return ans
}

//利用位掩码 0x55555555（二进制的 010101⋯），取出偶数下标比特和奇数下标比特，分别用库函数统计 1 的个数。

func evenOddBit(n int) []int {
	const mask = 0x5555
	return []int{bits.OnesCount16(uint16(n & mask)), bits.OnesCount16(uint16(n & (mask >> 1)))}
}
