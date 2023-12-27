package main

// 选哪个的问题？

func numOfBurgers(tomatoSlices int, cheeseSlices int) (ans []int) {
	// 这是2m
	m := (tomatoSlices - 2*cheeseSlices)
	n := (4*cheeseSlices - tomatoSlices)

	if m%2 != 0 || n%2 != 0 {
		return []int{}
	}
	return []int{m / 2, n / 2}
}
