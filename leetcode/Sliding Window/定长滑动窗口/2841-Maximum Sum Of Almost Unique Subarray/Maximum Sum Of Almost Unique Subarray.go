package main

func maxSum(nums []int, m int, k int) (ans int64) {
	n := len(nums)
	sum := 0
	hash := map[int]int{}
	for _, v := range nums[:k] {
		sum += v
		hash[v]++
	}
	if len(hash) >= m {
		ans = int64(sum)
	}
	for i := k; i < n; i++ {
		sum += nums[i] - nums[i-k]
		hash[nums[i]]++
		hash[nums[i-k]]--
		if hash[nums[i-k]] == 0 {
			delete(hash, nums[i-k])
		}
		if len(hash) >= m {
			ans = max(ans, int64(sum))
		}
	}
	return
}
