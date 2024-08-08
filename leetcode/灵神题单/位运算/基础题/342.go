package main

import "math/bits"

//  4    16
// 100  10000

func isPowerOfFour1(n int) bool {
	return bits.OnesCount(uint(n)) == 1 && bits.TrailingZeros(uint(n))&1 == 0
}

// mask=(10101010101010101010101010101010)_2

func isPowerOfFour2(n int) bool {
	const mask = 0xaaaaaaaa
	return n > 0 && n&(n-1) == 0 && n&mask == 0
}

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
