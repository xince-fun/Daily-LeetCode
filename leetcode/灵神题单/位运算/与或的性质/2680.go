package main

// 空间复杂度 O(n)
func maximumOr1(nums []int, k int) (ans int64) {
	n := len(nums)
	pre, suf := 0, make([]int, n+1)
	for i := 0; i < n; i++ {
		suf[n-i-1] = suf[n-i] | nums[n-i-1]
	}
	for i, x := range nums {
		num := int64(pre | suf[i+1] | (x << k))
		ans = max(ans, num)
		pre |= x
	}
	return
}
