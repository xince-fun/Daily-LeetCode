package main

import "fmt"

func countTestedDevices(batteryPercentages []int) int {
	ans := 0
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-ans > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	bat := []int{0, 1, 2}
	fmt.Println(countTestedDevices(bat))
}
