package main

func wateringPlants(plants []int, capacity int) int {
	ans := 0
	now := capacity
	i := 0
	for i < len(plants) {
		if now >= plants[i] {
			now -= plants[i]
			ans++
			i++
		} else {
			ans += i * 2
			now = capacity
		}
	}
	return ans
}
