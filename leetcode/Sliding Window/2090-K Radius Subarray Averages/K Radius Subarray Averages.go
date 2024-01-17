package main

func getAverages(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	sum := 0
	for i, v := range nums {
		if i < k || i+k >= n {
			ans[i] = -1
		}
		sum += v
		if i >= k*2 {
			ans[i-k] = sum / (k*2 + 1)
			sum -= nums[i-k*2]
		}
	}
	return ans
}
