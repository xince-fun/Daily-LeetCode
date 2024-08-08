package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k-1; i++ {
		sum += nums[i]
	}
	ans := float64(-10e4 - 1)
	for i := k - 1; i < len(nums); i++ {
		sum += nums[i]
		ans = max(ans, float64(sum)/float64(k))
		sum -= nums[i-k+1]
	}
	return ans
}
