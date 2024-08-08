package main

func maximumSubarraySum(nums []int, k int) int64 {
	count := map[int]int{}
	sum, ans := 0, 0

	for i, in := range nums {
		sum += in
		count[in]++

		if i < k-1 {
			continue
		}

		if len(count) == k {
			ans = max(ans, sum)
		}
		out := nums[i-k+1]
		count[out]--
		if count[out] == 0 {
			delete(count, out)
		}
		sum -= out
	}
	return int64(ans)
}
