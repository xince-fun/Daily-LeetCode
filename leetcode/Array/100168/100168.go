package main

func minOperations(nums []int, k int) int {
	currentXOR := 0
	for _, num := range nums {
		currentXOR ^= num
	}
	// 计算得到k需要的异或变化
	xorNeeded := currentXOR ^ k

	// 计算最少操作数，即xorNeeded中1的个数
	operations := 0
	for xorNeeded > 0 {
		// 翻转最低位的1
		operations += xorNeeded & 1
		// 右移一位
		xorNeeded >>= 1
	}

	return operations
}
