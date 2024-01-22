package main

import "fmt"

// 枚举相等的元素
func longestEqualSubarray1(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i-len(pos[x]))
	}
	for _, ps := range pos {
		if len(ps) <= ans {
			continue
		}
		left := 0
		for right, p := range ps {
			for p-ps[left] > k {
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}
	for _, ps := range pos {
		if len(ps) <= ans {
			continue
		}
		left := 0
		for right, p := range ps {
			for p-ps[left]-(right-left) > k {
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func main() {
	nums := []int{1, 3, 2, 3, 1, 3}
	k := 3
	fmt.Println(longestEqualSubarray(nums, k))
}
