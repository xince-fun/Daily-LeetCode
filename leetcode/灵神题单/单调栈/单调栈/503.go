package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for i := 0; i < 2*n; i++ {
		for len(st) > 0 && nums[i%n] > nums[st[len(st)-1]] {
			x := st[len(st)-1]
			st = st[:len(st)-1]
			ans[x] = nums[i%n]
		}
		st = append(st, i%n)
	}
	return ans
}
