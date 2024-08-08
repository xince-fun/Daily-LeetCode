package main

func maximumUniqueSubarray(nums []int) (ans int) {
	cnt := map[int]int{}
	sum, left := 0, 0
	for right, num := range nums {
		sum += num
		cnt[num]++
		for right-left+1 > len(cnt) {
			sum -= nums[left]
			cnt[nums[left]]--
			if cnt[nums[left]] == 0 {
				delete(cnt, nums[left])
			}
			left++
		}
		ans = max(ans, sum)
	}
	return
}
