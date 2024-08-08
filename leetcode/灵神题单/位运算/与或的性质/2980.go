package main

func hasTrailingZeros1(nums []int) bool {
	cnt := 0
	for _, num := range nums {
		if num&1 == 0 {
			cnt++
		}
		if cnt > 1 {
			return true
		}
	}
	return false
}

func hasTrailingZeros(nums []int) bool {
	cnt := len(nums)
	for _, x := range nums {
		cnt -= x & 1
	}
	return cnt >= 2
}
