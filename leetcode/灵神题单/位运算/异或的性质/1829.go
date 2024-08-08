package main

func getMaximumXor(nums []int, maximumBit int) (ans []int) {
	n := len(nums)
	mask := (1 << maximumBit) - 1
	xorsum := 0
	for _, x := range nums {
		xorsum ^= x
	}
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, xorsum^mask)
		xorsum ^= nums[i]
	}
	return
}
