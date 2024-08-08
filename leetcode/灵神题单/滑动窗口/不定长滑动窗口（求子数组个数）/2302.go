package main

func countSubarrays(nums []int, k int64) int64 {
	ans := 0
	total := 0
	left := 0
	for right, num := range nums {
		total += num

		for int64(total*(right-left+1)) >= k {
			total -= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return int64(ans)
}
