package main

// 按位与最大，就是找到数组中最大的值，连续的最大长度

func longestSubarray(nums []int) (ans int) {
	max, cnt := 0, 0
	for _, x := range nums {
		if x > max {
			max, ans, cnt = x, 1, 1
		} else if x < max {
			cnt = 0
		} else if cnt++; cnt > ans {
			ans = cnt
		}
	}
	return
}
