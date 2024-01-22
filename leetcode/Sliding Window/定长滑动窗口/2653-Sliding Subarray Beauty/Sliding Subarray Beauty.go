package main

func getSubarrayBeauty(nums []int, k int, x int) (ans []int) {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] {
		cnt[num+bias]++
	}
	ans = make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++
		left := x
		// 暴力枚举 [-50, -1]
		for j, c := range cnt[:bias] {
			left -= c
			if left <= 0 {
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]--
	}
	return
}
