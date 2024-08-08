package main

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
