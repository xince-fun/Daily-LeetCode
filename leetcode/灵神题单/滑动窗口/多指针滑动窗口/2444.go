package main

func countSubarrays(nums []int, minK int, maxK int) (ans int64) {
	minI, maxI, i0 := -1, -1, -1
	for right, num := range nums {
		if num == minK {
			minI = right
		}
		if num == maxK {
			maxI = right
		}
		if num < minK || num > maxK {
			i0 = right
		}
		ans += int64(max(min(minI, maxI)-i0, 0))
	}
	return
}
