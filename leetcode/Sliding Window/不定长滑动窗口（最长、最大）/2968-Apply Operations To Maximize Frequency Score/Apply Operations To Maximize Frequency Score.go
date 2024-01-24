package main

import "sort"

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
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
