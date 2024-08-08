package main

func xorOperation1(n int, start int) (ans int) {
	for i := 0; i < n; i++ {
		ans ^= start + 2*i
	}
	return
}

func xorN(n int) int {
	switch n % 4 {
	case 0:
		return n
	case 1:
		return 1
	case 2:
		return n + 1
	default:
		return 0
	}
}

func xorOperation(n int, start int) (ans int) {
	a := start / 2
	b := n & start & 1 // 都为奇数才为 1
	return (xorN(a+n-1)^xorN(a-1))*2 + b
}
