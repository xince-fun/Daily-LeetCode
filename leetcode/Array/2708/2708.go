package main

func maxStrength(nums []int) int64 {
	mn, mx := nums[0], nums[0]
	for _, x := range nums[1:] {
		mn, mx = min(mn, x, mn*x, mx*x),
			max(mx, x, mn*x, mx*x)
	}
	return int64(mx)
}
