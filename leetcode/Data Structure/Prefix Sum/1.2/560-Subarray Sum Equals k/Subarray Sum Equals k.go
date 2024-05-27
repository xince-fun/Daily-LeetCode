package main

func subarraySum1(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	m := make(map[int]int)
	ans := 0
	for i := 0; i < n; i++ {
		r := s[i+1]
		l := r - k
		ans += m[l]
		m[r]++
	}
	return ans
}

func subarraySum(nums []int, k int) (ans int) {
	n := len(nums)
	m := make(map[int]int)
	m[0] = 1
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		ans += m[sum-k]
		m[sum]++
	}
	return
}
