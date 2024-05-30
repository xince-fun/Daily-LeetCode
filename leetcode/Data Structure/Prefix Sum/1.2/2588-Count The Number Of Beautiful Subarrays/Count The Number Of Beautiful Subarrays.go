package main

func beautifulSubarrays1(nums []int) (ans int64) {
	n := len(nums)
	preSum := make([]int, n+1)
	for i, x := range nums {
		preSum[i+1] = preSum[i] ^ x
	}
	m := map[int]int{}

	for _, x := range preSum {
		ans += int64(m[x])
		m[x]++
	}
	return
}

func beautifulSubarrays(nums []int) (ans int64) {
	s := 0
	cnt := map[int]int{s: 1}
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
