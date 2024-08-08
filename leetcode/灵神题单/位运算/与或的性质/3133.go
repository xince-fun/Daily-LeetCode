package main

/*
1. 怎么保证 nums[i] 的 AND 是 x?
x 的第 i 个比特是 1，那么 nums[i] 的第 i 个比特也必须是 1



*/

func minEnd1(n int, x int) int64 {
	n--
	i, j := 0, 0
	// n 的 第 j 个比特（从右往左读） 填到 x 的第 i个比特位上
	for n>>j > 0 {
		if x>>i&1 == 0 {
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}

func minEnd(n int, x int) int64 {
	n--
	j := 0
	for t, lb := ^x, 0; n>>j > 0; t ^= lb {
		lb = t & -t
		x |= n >> j & 1 * lb
		j++
	}
	return int64(x)
}
