package main

func incremovableSubarrayCount(nums []int) (ans int) {
	n := len(nums)

	isIncresing := func(l, r int) bool {
		for i := 1; i < n; i++ {
			if i >= l && i <= r+1 {
				continue
			}
			if nums[i] <= nums[i-1] {
				return false
			}
		}
		if l-1 >= 0 && r+1 < n && nums[r+1] <= nums[l-1] {
			return false
		}
		return true
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isIncresing(i, j) {
				ans++
			}
		}
	}
	return
}
