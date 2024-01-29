package main

import "fmt"

func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		i0 := i
		for i++; i < n && nums[i] == nums[i-1]+1; i++ {
		}
		if i0 == i {
			ans = append(ans, fmt.Sprintf("%d", nums[i0]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[i0], nums[i-1]))
		}
	}
	return
}
