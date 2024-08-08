package main

import (
	"math"
	"slices"
	"sort"
)

func repairCars1(ranks []int, cars int) int64 {
	left := 0
	// r 4r 9r 16r 25r
	// 二分时间 t
	right := slices.Max(ranks)*cars*cars + 1

	check := func(mid int) bool {
		// 找到小于等于 mid 的第一个 n 车辆
		sum := 0
		for _, r := range ranks {
			sum += int(math.Sqrt(float64(mid / r)))
		}
		return sum >= cars
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

func repairCars(ranks []int, cars int) int64 {
	return int64(sort.Search(slices.Min(ranks)*cars*cars, func(i int) bool {
		sum := 0
		for _, r := range ranks {
			sum += int(math.Sqrt(float64(i / r)))
		}
		return sum >= cars
	}))
}
