package main

func temperatureTrend(temperatureA []int, temperatureB []int) (ans int) {
	n := len(temperatureA)
	lA, lB := temperatureA[0], temperatureB[0]
	iA := 0
	for i := 1; i < n; i++ {
		nowA, nowB := temperatureA[i]-lA, temperatureB[i]-lB
		if convert(nowA) == convert(nowB) {
			ans = max(ans, i-iA)
		} else {
			iA = i
		}
		lA, lB = temperatureA[i], temperatureB[i]
	}
	return
}

func convert(a int) int {
	if a == 0 {
		return 0
	} else if a > 0 {
		return 1
	}
	return -1
}
