package main

import (
	"slices"
)

// 前缀和 pre[i] 表示 i 之前的和 pre[0]=0
// pre[2] - pre[0] = nums[0] + nums[1]

// 此时以nums[right]作为最大频数
// 总sum = \sum(xr-xi) = (right-left)xr + \sum xi
// 每次移动的 (x_{right+1} - x_{right})*(right-left+1)

// 左移动 (x_{right+1}-x_{left})

func maxFrequency(nums []int, k int) (ans int) {
	slices.Sort(nums)
	left := 0
	sum := 0
	ans = 1
	for right := 1; right < len(nums); right++ {
		sum += (nums[right] - nums[right-1]) * (right - left)
		for sum > k {
			sum -= nums[right] - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
