package main

func maxSubarrayLength(nums []int, k int) (ans int) {
	left := 0
	freq := make(map[int]int)
	for right, x := range nums {
		freq[x]++
		for freq[x] > k {
			freq[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
