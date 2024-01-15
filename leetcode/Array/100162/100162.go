package main

func maxFrequencyElements(nums []int) int {
	m := make(map[int]int)
	maxFreq := 0

	for _, n := range nums {
		m[n]++
		if m[n] > maxFreq {
			maxFreq = m[n]
		}
	}
	total := 0
	for _, v := range m {
		if v == maxFreq {
			total += v
		}
	}
	return total
}
