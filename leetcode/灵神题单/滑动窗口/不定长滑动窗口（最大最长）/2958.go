package main

func maxSubarrayLength1(nums []int, k int) (ans int) {
	maxK, left := 0, 0
	cnt := map[int]int{}
	for right, num := range nums {
		cnt[num]++
		maxK = max(maxK, cnt[num])
		for maxK > k {
			out := nums[left]
			cnt[out]--
			if num == out {
				maxK--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range nums {
		cnt[x]++
		for cnt[x] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
