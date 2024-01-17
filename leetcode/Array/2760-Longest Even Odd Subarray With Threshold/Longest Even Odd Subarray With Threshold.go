package main

func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++
			continue
		}
		start := i
		i++
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-start)
	}
	return
}
