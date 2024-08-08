package main

func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 { // 有交集
			or ^= nums[left] // 从 or中去掉
			left++
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return
}
