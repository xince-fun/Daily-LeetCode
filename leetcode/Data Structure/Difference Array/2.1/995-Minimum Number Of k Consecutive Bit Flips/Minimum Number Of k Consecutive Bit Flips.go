package main

func minKBitFlips(nums []int, k int) (ans int) {
	n := len(nums)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range nums {
		revCnt += diff[i]
		if (v+revCnt)%2 == 0 {
			if i+k > n {
				return -1
			}
			ans++
			revCnt++
			diff[i+k]--
		}
	}
	return
}

func minKBitFlips1(nums []int, k int) (ans int) {
	n := len(nums)
	diff := make([]int, n+1)
	revCnt := 0

	for i, v := range nums {
		revCnt ^= diff[i]
		if v == revCnt {
			if i+k > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+k] ^= 1
		}
	}
	return
}
