package main

func maximumUniqueSubarray(nums []int) (ans int) {
	left, sum := 0, 0
	hash := make(map[int]int)
	for _, x := range nums {
		hash[x]++
		sum += x
		for hash[x] > 1 {
			sum -= nums[left]
			hash[nums[left]]--
			if hash[nums[left]] == 0 {
				delete(hash, nums[left])
			}
			left++
		}
		ans = max(ans, sum)
	}
	return
}
