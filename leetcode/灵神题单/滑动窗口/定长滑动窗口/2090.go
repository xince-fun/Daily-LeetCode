package main

func getAverages1(nums []int, k int) []int {
	sum := 0
	ans := make([]int, len(nums))
	for i := 0; i < k && i < len(nums); i++ {
		sum += nums[i]
		ans[i] = -1
	}
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		if i >= 2*k {
			ans[i-k] = sum / (2*k + 1)
			sum -= nums[i-2*k]
		}
	}
	for i := len(nums) - 1; i > len(nums)-1-k && i > 0; i-- {
		ans[i] = -1
	}
	return ans
}

func getAverages(nums []int, k int) []int {
	sum := 0
	ans := make([]int, len(nums))
	for i, v := range nums {
		if i < k || i+k >= len(nums) {
			ans[i] = -1
		}
		sum += v
		if i >= 2*k {
			ans[i-k] = sum / (2*k + 1)
			sum -= nums[i-2*k]
		}
	}
	return ans
}
