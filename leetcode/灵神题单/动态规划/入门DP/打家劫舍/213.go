package main

func rob1(nums []int, start, end int) int {
	f0, f1 := 0, 0
	for i := start; i < end; i++ {
		f0, f1 = f1, max(f1, f0+nums[i])
	}
	return f1
}

// 考虑偷不偷 nums[0]

// 偷nums[0] nums[1]和nums[n-1]不能偷 变成 nums[2]到nums[n-2]
// 不偷nums[0] nums[1]到nums[n-1]的问题

func rob(nums []int) int {
	n := len(nums)
	return max(nums[0]+rob1(nums, 2, n-1), rob1(nums, 1, n))
}
