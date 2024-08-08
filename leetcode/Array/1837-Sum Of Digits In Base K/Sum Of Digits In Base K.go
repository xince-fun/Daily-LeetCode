package main

func sumBase(n int, k int) (ans int) {
	// 除余法
	for n > 0 {
		ans += n % k
		n /= k
	}
	return
}
