package main

import "sort"

/*

一. 前缀和做法
1. 改的是排序后的一段连续子数组
2. 使用滑动窗口
3.

二. 不使用前缀和做法
1. 如果数组长度是奇数 n
   距离和 = a[n-1] + a[n-2] + .. + a[n/2+1]
         - (a[0] + a[1] + ... + a[n/2-1])
   如果数组长度是偶数 n
   距离和 = a[n-1] + a[n-2] + .. + a[n/2]
         - (a[0] + a[1] + ... + a[n/2-1])

2. 随着窗口向右移动，每个数字会经历三个过程
	1. 移入窗口，对距离和的贡献是 + nums[i]
	2. 变成正中间的数字，对距离和的贡献是 0
	3. 变成左半边的数字，对距离和的贡献是 - nums[i]

	从 1->2 和从 2->3 都是发生在左边的中位数上
	从 3->2 和从 2->1 都是发生在右边的中位数上

*/

func maxFrequencyScore(nums []int, k int64) int {
	sort.Ints(nums)
	n := len(nums)
	// 长为n+1的前缀和数组
	s := make([]int, n+1)
	for i, v := range nums {
		s[i+1] = s[i] + v
	}

	distanceSum := func(left, right int) int64 {
		i := (left + right) / 2
		leftSum := nums[i]*(i-left) - (s[i] - s[left])
		rightSum := s[right+1] - s[i+1] - nums[i]*(right-i)
		return int64(leftSum) + int64(rightSum)
	}

	ans, left := 0, 0
	for right := 0; right < n; right++ {
		for distanceSum(left, right) > k {
			left += 1
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFrequencyScore1(nums []int, k int64) int {
	sort.Ints(nums)
	n := len(nums)

	ans, left := 0, 0
	s := int64(0) // 到中位数的距离和
	for right := 0; right < n; right++ {
		s += int64(nums[right] - nums[(left+right)/2])
		for s > k {
			s += int64(nums[left] - nums[(left+right+1)/2])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
