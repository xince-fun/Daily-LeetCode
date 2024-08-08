package main

const mod = 1_000_000_007

func maxSubArray(nums []int) (ans, sum int) {
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
		sum += x
	}
	return
}

func kConcatenationMaxSum(arr []int, k int) int {
	ans, sum := maxSubArray(arr)
	if k == 1 {
		return ans
	}
	arr = append(arr, arr...)
	ans, _ = maxSubArray(arr)
	return max(ans, sum*(k-2)+ans) % mod
}
