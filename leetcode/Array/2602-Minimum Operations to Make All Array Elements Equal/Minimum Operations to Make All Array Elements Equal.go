package main

import "sort"

func minOperations(nums []int, queries []int) (ans []int64) {
	sort.Ints(nums)
	n := len(nums)
	// 长度为n+1的前缀和
	s := make([]int, n+1)
	for i, v := range nums {
		s[i+1] = s[i] + v
	}

	for _, q := range queries {
		i := sort.SearchInts(nums, q)
		leftSum := q*i - s[i]
		rightSum := s[n] - s[i+1] - q*(n-i)
		ans = append(ans, int64(leftSum)+int64(rightSum))
	}
	return
}

func main() {

}
