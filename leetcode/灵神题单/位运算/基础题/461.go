package main

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func hammingDistance1(x int, y int) int {
	return onesCount(x ^ y)
}

func hammingDistance2(x, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}

func hammingDistance(x, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}
