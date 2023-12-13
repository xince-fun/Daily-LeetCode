package main

func nextGreaterElements(nums []int) []int {
	// 这里是两倍元素
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := make([]int, 0)
	for i := 0; i < n*2; i++ {
		for len(st) != 0 && nums[i%n] > nums[st[len(st)-1]] {
			x := st[len(st)-1]
			st = st[:len(st)-1]
			ans[x] = nums[i%n]
		}
		st = append(st, i%n)
	}
	return ans
}
