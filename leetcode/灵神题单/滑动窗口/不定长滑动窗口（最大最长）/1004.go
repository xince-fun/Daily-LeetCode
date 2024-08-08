package main

func longestOnes1(nums []int, k int) (ans int) {
	left := 0
	cnt := 0
	for right, num := range nums {
		cnt += num
		for right-left+1 > cnt+k {
			cnt -= nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

// 统计0的个数
func longestOnes(nums []int, k int) (ans int) {
	left, cnt := 0, 0
	for right, x := range nums {
		cnt += x ^ 1
		for cnt > k {
			cnt -= nums[left] ^ 1
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
