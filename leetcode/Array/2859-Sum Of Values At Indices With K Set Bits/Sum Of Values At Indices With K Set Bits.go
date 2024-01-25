package main

import (
	"fmt"
	"math/bits"
)

func sumIndicesWithKSetBits(nums []int, k int) (sum int) {
	for i, num := range nums {
		if bits.OnesCount(uint(i)) == k {
			sum += num
		}
	}
	return
}

func main() {
	nums := []int{5, 10, 1, 5, 2}
	k := 1
	fmt.Println(sumIndicesWithKSetBits(nums, k))
}
