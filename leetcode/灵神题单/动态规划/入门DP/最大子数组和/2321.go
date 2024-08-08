package main

// nums1 = [60,60,90] nums2 = [10,90,10]

// sum = s_1 - (nums1[left] + ... + nums1[right]) + (nums2[left] + ... + nums2[right])

// s1 + (nums2[left] - nums1[left]) + .. (nums2[right]- nums2[left])

func solve(nums1, nums2 []int) int {
	s1, maxSum, s := 0, 0, 0
	for i, x := range nums1 {
		s1 += x
		s = max(s+nums2[i]-x, 0)
		maxSum = max(maxSum, s)
	}
	return s1 + maxSum
}

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1))
}
