package main

/*
1. 如果 nums 的最后两个数相等，去掉这两个数，问题变成剩下的n-2个数
2. 如果 nums 的最后三个数相等，那么去掉这三个数，问题变成剩下的n-3个数
3. 如果 nums 的最后三个数是连续递增的，那么去掉这三个数，问题变成剩下的n-3个数

f[0] = true, f[i+1]表示能否有效划分 nums[0] 到 nums[i]

		     f[i-1] ^ nums[i] = nums[i-1]						 i > 0
f[i+1] = V   f[i-2] ^ nums[i] = nums[i-1] = nums[i-2]			 i > 1
		     f[i-2] ^ nums[i] = nums[i-1] + 1 = nums[i-2] + 2    i > 1
*/

func validPartition(nums []int) bool {
	// 这是个 dp 问题
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}
