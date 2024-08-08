package main

import (
	"slices"
)

func binarySearch(arr []int, low, high int) bool {
	left, right := -1, len(arr)
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] < low {
			left = mid
		} else if arr[mid] > high {
			right = mid
		} else if arr[mid] >= low && arr[mid] <= high {
			return true
		}
	}
	return false
}

// -d + y <= x <= d + y

func findTheDistanceValue(arr1 []int, arr2 []int, d int) (ans int) {
	slices.Sort(arr2)
	for _, x := range arr1 {
		low, high := x-d, x+d
		if !binarySearch(arr2, low, high) {
			ans++
		}
	}
	return ans
}
