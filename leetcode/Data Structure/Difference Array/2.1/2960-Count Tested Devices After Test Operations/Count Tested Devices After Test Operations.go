package main

func countTestedDevices(batteryPercentages []int) int {
	dec := 0
	for _, x := range batteryPercentages {
		if x > dec {
			dec++
		}
	}
	return dec
}
