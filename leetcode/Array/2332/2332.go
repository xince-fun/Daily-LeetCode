package main

import (
	"fmt"
	"slices"
)

// buses = [10,20,30]
// passengers = [4,11,13,19,21,24,26]

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) (ans int) {
	slices.Sort(buses)
	slices.Sort(passengers)

	j, c := 0, 0

	for _, t := range buses {
		for c = capacity; c > 0 && j < len(passengers) && passengers[j] <= t; c-- {
			j++
		}
	}

	if c > 0 {
		ans = buses[len(buses)-1]
	} else {
		ans = passengers[j-1]
	}
	for j--; j >= 0 && ans == passengers[j]; j-- {
		ans--
	}

	return
}

func main() {
	// buses := []int{10, 20}
	buses := []int{20, 30, 10}
	// passengers := []int{2, 17, 18, 19}
	passengers := []int{19, 13, 26, 4, 25, 11, 21}
	capacity := 2
	fmt.Println(latestTimeCatchTheBus(buses, passengers, capacity))
}
