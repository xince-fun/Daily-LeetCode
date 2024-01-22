package main

import "fmt"

func minSwaps(nums []int) (ans int) {
	k := 0
	for _, num := range nums {
		k += num
	}
	nums = append(nums, nums...)
	cnt, maxCnt := 0, 0
	for i, num := range nums {
		cnt += num
		if i >= k {
			cnt -= nums[i-k]
			if cnt > maxCnt {
				maxCnt = cnt
			}
		}
	}
	return k - maxCnt
}

func main() {
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	fmt.Println(minSwaps(nums))
}
