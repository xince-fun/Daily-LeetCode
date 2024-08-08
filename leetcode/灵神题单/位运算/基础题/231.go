package main

import "math/bits"

//  2  4  8
// 10 100 1000

func isPowerOfTwo1(n int) bool {
	return bits.OnesCount(uint(n)) == 1
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
