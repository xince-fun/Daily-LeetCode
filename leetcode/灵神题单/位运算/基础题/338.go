package main

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits1(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = onesCount(i)
	}
	return ans
}

// i&(i-1) = 0, 更新 highBit=i, 表示当前的最高有效位
// i 比 i-highBit 的 「一比特数」 多 1，
// bits[i] = bits[i-highBit] + 1

func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

// bits[x] = bits[x>>1] + (x&1)

func countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + (i & 1)
	}
	return bits
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
