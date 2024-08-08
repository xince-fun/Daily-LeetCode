package main

import (
	"slices"
	"sort"
)

func minimumTime1(time []int, totalTrips int) int64 {
	left, right := 0, totalTrips*1e7
	check := func(t int) bool {
		sum := 0
		for _, tt := range time {
			sum += t / tt
		}
		return sum >= totalTrips
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return int64(right)
}

func minimumTime(time []int, totalTrips int) int64 {
	minT := slices.Min(time)
	left, right := minT-1, totalTrips*minT
	for left+1 < right {
		mid := left + (right-left)/2
		sum := 0
		for _, t := range time {
			sum += mid / t
		}
		if sum >= totalTrips {
			right = mid
		} else {
			left = mid
		}
	}
	return int64(right)
}

func minimumTime2(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(i int) bool {
		check := func(t int) bool {
			sum := 0
			for _, tt := range time {
				sum += t / tt
			}
			return sum >= totalTrips
		}
		return check(i)
	}))
}
