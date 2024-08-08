package main

// 暴力
func maxWidthRamp1(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if nums[j] >= nums[i] {
				ans = max(ans, j-i)
				break
			}
		}
	}
	return
}

// 以 nums[0] 为基准的严格单调递减栈

// 当遍历到 i_0时， 若存在 i_0 < i_1, 且 nums[i_0] <= nums[i_1]
// 若 nums[j] < nums[i_1] 不可能是答案
// 若 nums[j] >= nums[i_1]， nums[j] >= nums[i_0], 那么 (i_0,j)优于 (i_1,j)

// 所以遍历到 i_0时， 若存在 i_0 < i_1, 且 nums[i_0] <= nums[i_1], 不用考虑 i_1

// 从 下标 0 存进 st 开始， nums[st[1]] < nums[st[0]], nums[st[2]] < nums[st[1]]
// st 就是一个严格单调递减栈

func maxWidthRamp2(nums []int) (ans int) {
	st := []int{}
	for i, x := range nums {
		if len(st) == 0 || x < nums[st[len(st)-1]] {
			st = append(st, i)
		}
	}
	j := len(nums) - 1
	for len(st) > 0 && j >= 0 {
		for j >= 0 && nums[j] < nums[st[len(st)-1]] {
			j--
		}
		if j >= 0 {
			ans = max(ans, j-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return
}

func maxWidthRamp(nums []int) (ans int) {
	st := []int{0}
	for j, x := range nums {
		if x < nums[st[len(st)-1]] {
			st = append(st, j)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] >= nums[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return
}
