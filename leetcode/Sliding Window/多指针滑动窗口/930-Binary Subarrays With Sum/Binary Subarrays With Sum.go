package main

import (
	"fmt"
)

/*
计算 nums 的前缀和 sum
*/

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < n; i++ {
		r := sum[i+1]
		l := r - goal
		ans += m[l]
		m[r]++
	}
	return
}

/*
nums[i] 没有负权值意味着前缀和数组必然具有（非严格）单调递增特性。

我们可以使用两个左端点 l1和 l2，代表在给定右端点 r 的前提下满足要求的左端点集合，
同时使用 s1 和 s2 分别代表两个端点到 r 这一段的和。

*/

func numSubarraysWithSum1(nums []int, goal int) (ans int) {
	n := len(nums)
	for l1, l2, r, s1, s2 := 0, 0, 0, 0, 0; r < n; r++ {
		s1 += nums[r]
		s2 += nums[r]
		for l1 <= r && s1 > goal {
			s1 -= nums[l1]
			l1++
		}
		for l2 <= r && s2 >= goal {
			s2 -= nums[l2]
			l2++
		}
		ans += l2 - l1
	}
	return
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	fmt.Println(numSubarraysWithSum(nums, goal))
}
