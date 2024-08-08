package main

// 1 1 1 1 1
// C52 = 5*4/2=10
// C42 = 6

// Cn2 - C(n-1)2 = n*(n-1)/2 - (n-1)(n-2)/2 = (n-1) = 4

// C22=1

func countGood(nums []int, k int) int64 {
	cnt := map[int]int{}
	ans, left := int64(0), 0
	total := 0
	for right, num := range nums {
		cnt[num]++
		if cnt[num] > 1 {
			total += cnt[num] - 1
		}
		for total >= k {
			out := nums[left]
			total -= cnt[out]
			cnt[out]--
			left++
		}
		ans += int64(right - left + 1)
	}
	return ans
}
