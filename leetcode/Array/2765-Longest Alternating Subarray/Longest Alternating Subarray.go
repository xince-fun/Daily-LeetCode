package main

func alternatingSubarray(nums []int) int {
	ans := -1
	i, n := 0, len(nums)
	for i < n-1 {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}
		i0 := i // 记录这一组的开始位置
		i += 2
		for i < n && nums[i] == nums[i0]+(i-i0)%2 {
			i++
		}
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
