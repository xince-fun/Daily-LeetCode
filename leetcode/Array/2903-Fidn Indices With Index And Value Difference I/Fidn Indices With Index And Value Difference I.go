package main

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	ans := []int{-1, -1}
	for i := 0; i < len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				ans[0], ans[1] = i, j
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
