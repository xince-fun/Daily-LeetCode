package main

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < n; i++ {
		r := s[i+1]
		l := r - goal
		ans += m[l]
		m[r]++
	}
	return
}
