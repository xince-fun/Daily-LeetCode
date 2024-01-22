package main

// 剩下的元素值应该是 sum - x
// 找到一个最长的序列，和为sum - x
func minOperations(nums []int, x int) (ans int) {
	left := 0
	sum := 0
	for _, num := range nums {
		sum += num
	}
	sum -= x
	if sum < 0 {
		return -1
	}
	ans = -1
	total := 0
	for right, num := range nums {
		total += num
		for total > sum {
			total -= nums[left]
			left++
		}
		if total == sum {
			ans = max(ans, right-left+1)
		}
	}
	if ans == -1 {
		return -1
	}
	return len(nums) - ans
}
