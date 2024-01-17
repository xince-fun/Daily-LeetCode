package main

import "fmt"

func maxSubarrayLength(nums []int, k int) int {
	ans, left := 0, 0
	m := make(map[int]int)
	for right := 0; right < len(nums); right++ {
		m[nums[right]]++
		// 如果超过k了
		if m[nums[right]] > k {
			for m[nums[right]] > k {
				l := nums[left]
				m[l]--
				left++
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
	k := 2
	fmt.Println(maxSubarrayLength(nums, k))
}
