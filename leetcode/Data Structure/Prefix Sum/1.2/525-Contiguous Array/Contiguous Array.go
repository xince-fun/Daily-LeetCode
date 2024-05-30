package main

func findMaxLength(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	m := map[int]int{}
	sum := 0
	for i, x := range nums {
		sum += x
		if sum == 0 {
			ans = i + 1
		}
		if m[sum] > 0 {
			ans = max(ans, i-m[sum])
		} else {
			m[sum] = i
		}
	}
	return
}
