package main

import "sort"

type MountainArray struct{}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

// [1,2,3,4,5,3,1]

// nums[mid] > target

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	peak := sort.Search(n-1, func(i int) bool { return mountainArr.get(i) >= mountainArr.get(i+1) })
	i := sort.Search(peak, func(i int) bool { return mountainArr.get(i) >= target })
	if mountainArr.get(i) != target {
		i = peak + sort.Search(n-1-peak, func(i int) bool { return mountainArr.get(peak+i) <= target })
		if mountainArr.get(i) != target {
			return -1
		}
	}
	return i
}
