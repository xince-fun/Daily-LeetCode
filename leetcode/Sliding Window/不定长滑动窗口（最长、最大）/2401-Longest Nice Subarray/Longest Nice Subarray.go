package main

func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 { // 有交集
			or ^= nums[left]
			left++
		}
		or |= x // 把集合 x 并入 or
		ans = max(ans, right-left+1)
	}
	return
}
