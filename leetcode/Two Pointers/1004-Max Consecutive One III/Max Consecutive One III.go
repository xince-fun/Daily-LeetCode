package main

func longestOnes(nums []int, k int) int {
	ans := 0
	left, zeroCnt := 0, 0
	for right, x := range nums {
		if x == 0 {
			zeroCnt++
		}

		for zeroCnt > k {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
