package main

func countCompleteSubarrays(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	left := 0
	c := map[int]int{}
	for _, x := range nums {
		c[x]++
		for len(c) == len(cnt) {
			out := nums[left]
			c[out]--
			if c[out] == 0 {
				delete(c, out)
			}
			left++
		}
		ans += left
	}
	return
}
