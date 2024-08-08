package main

func hIndex(citations []int) int {
	left, right := 0, len(citations)+1
	for left+1 < right {
		mid := left + (right-left)/2
		if citations[len(citations)-mid] >= mid {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
