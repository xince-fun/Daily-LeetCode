package main

/*
怎么枚举
怎么把子数组分类？

考虑移除的子数组最后一个数是 a[n-1]

从右往左枚举 j，表示j-1是移除的子数组的右端点
维护 i，表示 i+1 是移除的子数组的左断电的最大值（在右端点固定为 j-1 的情况下）
那么把 i + 2 加到答案中，表示右端点固定为 j-1的情况下，有i+2个左端点

移动i的方式就是不断的左移i直到a[i] < a[j] 或者 i 是 -1

*/

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	var ans int64
	i := 0
	for i < n-1 && nums[i] < nums[i+1] {
		i += 1
	}
	if i == n-1 {
		return int64(n * (n + 1) / 2)
	}
	ans = int64(i + 2)
	j := n - 1
	for j == n-1 || nums[j] < nums[j+1] {
		for i >= 0 && nums[i] >= nums[j] {
			i -= 1
		}
		ans += int64(i + 2)
		j -= 1
	}
	return ans
}
