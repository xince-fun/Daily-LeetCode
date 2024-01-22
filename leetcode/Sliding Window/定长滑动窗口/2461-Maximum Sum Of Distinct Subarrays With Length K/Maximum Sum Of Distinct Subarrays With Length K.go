package main

func maximumSubarraySum(nums []int, k int) (ans int64) {
	sum := 0
	hash := map[int]int{}
	for _, v := range nums[:k] {
		sum += v
		hash[v]++
	}
	if len(hash) == k {
		ans = int64(sum)
	}
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		hash[nums[i]]++
		hash[nums[i-k]]--
		if hash[nums[i-k]] == 0 {
			delete(hash, nums[i-k])
		}
		if len(hash) == k {
			ans = max(ans, int64(sum))
		}
	}
	return
}
