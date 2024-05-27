package main

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans := 0
	nowA, nowB := capacityA, capacityB
	i, j := 0, len(plants)-1
	for i < j {
		if nowA >= plants[i] {
			nowA -= plants[i]
		} else {
			ans++
			nowA = capacityA - plants[i]
		}
		if nowB >= plants[j] {
			nowB -= plants[j]
		} else {
			ans++
			nowB = capacityB - plants[j]
		}
		i++
		j--
	}

	if i == j && !((nowA >= nowB && nowA >= plants[i]) || (nowB > nowA && nowB >= plants[j])) {
		ans++
	}
	return ans
}
