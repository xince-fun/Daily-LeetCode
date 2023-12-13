package main

import "fmt"

func powMod(a, b, mod int) int {
	result := 1
	base := a % mod

	for b > 0 {
		if b%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		b /= 2
	}

	return result
}

func getGoodIndices(variables [][]int, target int) []int {
	ans := []int{}
	for i, v := range variables {
		a, b, c, m := v[0], v[1], v[2], v[3]
		firstStep := powMod(a, b, 10)
		result := powMod(firstStep, c, m)
		if result == target {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	vari := [][]int{{2, 2, 3, 2}, {2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}}
	target := 0
	fmt.Println(getGoodIndices(vari, target))
}
