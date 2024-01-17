package main

import "fmt"

func countCompleteSubarrays(nums []int) (ans int) {
	hash := map[int]bool{}
	for _, n := range nums {
		hash[n] = true
	}
	m := len(hash)
	left := 0
	h := map[int]int{}
	for _, v := range nums {
		h[v]++
		for len(h) == m {
			x := nums[left]
			h[x]--
			if h[x] == 0 {
				delete(h, x)
			}
			left++
		}
		ans += left // 子数组左端点 < left 的都是合法的
	}
	return
}

func main() {
	nums := []int{1, 3, 1, 2, 2}
	fmt.Println(countCompleteSubarrays(nums))
}
