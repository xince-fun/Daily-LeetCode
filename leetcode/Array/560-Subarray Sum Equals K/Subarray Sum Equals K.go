package main

func subarraySum(nums []int, k int) (ans int) {
	n := len(nums)
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < n; i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			ans += hash[preSum-k]
		}
		hash[preSum]++
	}
	return
}
