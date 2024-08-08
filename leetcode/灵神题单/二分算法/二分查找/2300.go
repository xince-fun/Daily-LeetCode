package main

import (
	"slices"
	"sort"
)

// x * pos[i] >= success

func lowerBound(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func successfulPairs1(spells []int, potions []int, success int64) (ans []int) {
	slices.Sort(potions)
	for _, x := range spells {
		idx := lowerBound(potions, (int(success)-1)/x+1)
		ans = append(ans, len(potions)-idx)
	}
	return ans
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}
